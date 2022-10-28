package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

//go:generate mockgen -destination=./individual_mock.go -package=db . IndividualRepo

type IndividualRepo interface {
	GetAll(ctx context.Context, options api.GetAllOptions) ([]*api.Individual, error)
	GetByID(ctx context.Context, id string) (*api.Individual, error)
	Put(ctx context.Context, individual *api.Individual, fields []string) (*api.Individual, error)
	PutMany(ctx context.Context, individuals []*api.Individual, fields []string) ([]*api.Individual, error)
	SoftDelete(ctx context.Context, id string, countryId string) error
	SoftDeleteMany(ctx context.Context, ids []string, countryId string) error
}

type individualRepo struct {
	db *sqlx.DB
}

func (i individualRepo) driverName() string {
	d := i.db.DriverName()
	if d == "sqlite3" {
		return "sqlite"
	} else if d == "postgres" {
		return "postgres"
	} else {
		panic("unsupported driver")
	}
}

func (i individualRepo) GetAll(ctx context.Context, options api.GetAllOptions) ([]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.getAllInternal(ctx, tx, options)
	})
	if err != nil {

		return nil, err
	}
	return ret.([]*api.Individual), nil
}

func (i individualRepo) getAllInternal(ctx context.Context, tx *sqlx.Tx, options api.GetAllOptions) ([]*api.Individual, error) {

	l := logging.NewLogger(ctx)
	l.Debug("getting all individuals", zap.Any("options", options))

	var ret []*api.Individual
	args := []interface{}{}
	query := "SELECT * FROM individuals"

	var whereClauses []string

	var nextArg = func(arg interface{}) string {
		args = append(args, arg)
		return fmt.Sprintf("$%d", len(args))
	}

	// we do not want to return deleted individuals
	whereClauses = append(whereClauses, "deleted_at IS NULL")

	if len(options.FullName) != 0 {
		if i.driverName() == "sqlite" {
			whereClauses = append(whereClauses, "full_name LIKE "+nextArg("%"+options.FullName+"%")+" OR preferred_name LIKE "+nextArg("%"+options.FullName+"%"))
		} else if i.driverName() == "postgres" {
			whereClauses = append(whereClauses, "full_name ILIKE "+nextArg("%"+options.FullName+"%")+" OR preferred_name ILIKE "+nextArg("%"+options.FullName+"%"))
		}
	}

	if len(options.Address) != 0 {
		if i.driverName() == "sqlite" {
			whereClause := "address LIKE " + nextArg("%"+options.Address+"%")
			whereClauses = append(whereClauses, whereClause)
		} else if i.driverName() == "postgres" {
			whereClause := "address ILIKE " + nextArg("%"+options.Address+"%")
			whereClauses = append(whereClauses, whereClause)
		}
	}

	if len(options.Genders) != 0 {
		whereClause := "gender in ("
		for i, g := range options.Genders {
			if i != 0 {
				whereClause += ","
			}
			whereClause += nextArg(g)
		}
		whereClause += ")"
		whereClauses = append(whereClauses, whereClause)
	}

	if options.BirthDateFrom != nil {
		whereClauses = append(whereClauses, "birth_date >= "+nextArg(options.BirthDateFrom))
	}

	if options.BirthDateTo != nil {
		whereClauses = append(whereClauses, "birth_date <= "+nextArg(options.BirthDateTo))
	}

	if len(options.PhoneNumber) != 0 {
		normalizedPhoneNumber := api.NormalizePhoneNumber(options.PhoneNumber)
		whereClauses = append(whereClauses, "normalized_phone_number LIKE "+nextArg("%"+normalizedPhoneNumber+"%"))
	}

	if len(options.Email) != 0 {
		normalizedEmail := strings.ToLower(options.Email)
		whereClauses = append(whereClauses, "email = "+nextArg(normalizedEmail))
	}

	if options.IsMinor != nil {
		if options.IsMinorSelected() {
			whereClauses = append(whereClauses, "is_minor = "+nextArg(true))
		} else if options.IsNotMinorSelected() {
			whereClauses = append(whereClauses, "is_minor = "+nextArg(false))
		}
	}

	if options.PresentsProtectionConcerns != nil {
		if options.IsPresentsProtectionConcernsSelected() {
			whereClauses = append(whereClauses, "presents_protection_concerns = "+nextArg(true))
		} else if options.IsNotPresentsProtectionConcernsSelected() {
			whereClauses = append(whereClauses, "presents_protection_concerns = "+nextArg(false))
		}
	}

	if len(options.CountryID) != 0 {
		whereClauses = append(whereClauses, "country_id = "+nextArg(options.CountryID))
	}

	if len(options.DisplacementStatuses) != 0 {
		qry := "displacement_status in ("
		for _, ds := range options.DisplacementStatuses {
			args = append(args, ds)
			qry += fmt.Sprintf("$%d,", len(args))
		}
		qry = qry[:len(qry)-1]
		qry += ")"
		whereClauses = append(whereClauses, qry)
	}

	if len(whereClauses) != 0 {
		query = query + " WHERE "
		for i, whereClause := range whereClauses {
			if i != 0 {
				query = query + " AND "
			}
			query = query + "( " + whereClause + " )"
		}
	}

	query = query + " ORDER BY ID"
	if options.Take != 0 {
		query += fmt.Sprintf(" LIMIT %d", options.Take)
	}
	if options.Skip != 0 {
		query += fmt.Sprintf(fmt.Sprintf(" OFFSET %d", options.Skip))
	}
	err := tx.SelectContext(ctx, &ret, query, args...)
	if err != nil {
		l.Error("failed to get all individuals", zap.Error(err))
		return nil, err
	}
	return ret, err
}

func (i individualRepo) GetByID(ctx context.Context, id string) (*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.getByIdInternal(ctx, tx, id)
	})
	if err != nil {
		return nil, err
	}
	return ret.(*api.Individual), nil
}

func (i individualRepo) getByIdInternal(ctx context.Context, tx *sqlx.Tx, id string) (*api.Individual, error) {
	l := logging.NewLogger(ctx).With(zap.String("individual_id", id))
	l.Debug("getting individual by id")
	var ret = api.Individual{}
	err := tx.GetContext(ctx, &ret, "SELECT * FROM individuals WHERE id = $1 and deleted_at IS NULL", id)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (i individualRepo) PutMany(ctx context.Context, individuals []*api.Individual, fields []string) ([]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.putManyInternal(ctx, tx, individuals, fields)
	})
	if err != nil {
		return nil, err
	}
	return ret.([]*api.Individual), nil
}

func (i individualRepo) putManyInternal(ctx context.Context, tx *sqlx.Tx, individuals []*api.Individual, fields []string) ([]*api.Individual, error) {

	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)

	fieldsSet := containers.NewStringSet(fields...)
	if fieldsSet.Contains("phone_number") {
		fieldsSet.Add("normalized_phone_number")
	}
	fieldsSet.Remove("deleted_at")
	fieldsSet.Remove("created_at")
	fieldsSet.Remove("updated_at")
	fieldsSet.Add("id")
	fields = fieldsSet.Items()

	ret := make([]*api.Individual, 0, len(individuals))
	if err := batch(maxParams/len(fields), individuals, func(individualsInBatch []*api.Individual) error {
		args := make([]interface{}, 0)
		b := &strings.Builder{}
		b.WriteString("INSERT INTO individuals (" + strings.Join(fields, ",") + ",created_at,updated_at) VALUES ")

		for i, individual := range individualsInBatch {
			if individual.ID == "" {
				individual.ID = xid.New().String()
			}
			if i != 0 {
				b.WriteString(",")
			}
			b.WriteString("(")
			for j, field := range fields {
				if j != 0 {
					b.WriteString(",")
				}
				fieldValue, err := individual.GetFieldValue(field)
				if err != nil {
					return err
				}
				args = append(args, fieldValue)
				b.WriteString(fmt.Sprintf("$%d", len(args)))
			}
			b.WriteString(",'" + nowStr + "','" + nowStr + "')")
		}
		b.WriteString(" ON CONFLICT (id) DO UPDATE SET ")
		isFirst := true
		for _, field := range fields {
			if field == "id" {
				continue
			}
			if !isFirst {
				b.WriteString(",")
			}
			isFirst = false
			b.WriteString(fmt.Sprintf("%s = EXCLUDED.%s", field, field))
		}
		b.WriteString(", updated_at='" + nowStr + "'")
		b.WriteString(" RETURNING *")

		var out []*api.Individual
		qry := b.String()
		err := tx.SelectContext(ctx, &out, qry, args...)
		if err != nil {
			return err
		}
		ret = append(ret, out...)
		return nil
	}); err != nil {
		return nil, err
	}

	return ret, nil

}

func (i individualRepo) Put(ctx context.Context, individual *api.Individual, fields []string) (*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.putInternal(ctx, tx, individual, fields)
	})
	if err != nil {
		return nil, err
	}
	return ret.(*api.Individual), nil
}

func (i individualRepo) putInternal(ctx context.Context, tx *sqlx.Tx, individual *api.Individual, fields []string) (*api.Individual, error) {
	ret, err := i.putManyInternal(ctx, tx, []*api.Individual{individual}, fields)
	if err != nil {
		return nil, err
	}
	if len(ret) != 1 {
		return nil, fmt.Errorf("unexpected number of individuals returned: %d", len(ret))
	}
	return ret[0], nil
}

func (i individualRepo) SoftDelete(ctx context.Context, id string, countryId string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.softDeleteManyInternal(ctx, tx, []string{id}, countryId)
		return nil, err
	})
	return err
}

func (i individualRepo) SoftDeleteMany(ctx context.Context, ids []string, countryId string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.softDeleteManyInternal(ctx, tx, ids, countryId)
		return nil, err
	})
	return err
}

func (i individualRepo) softDeleteManyInternal(ctx context.Context, tx *sqlx.Tx, ids []string, countryId string) error {
	idSet := containers.NewStringSet(ids...)
	ids = idSet.Items()

	l := logging.NewLogger(ctx).With(zap.Strings("individual_ids", ids))
	l.Debug("deleting individuals")

	const query = "UPDATE individuals SET deleted_at = $1 WHERE id IN ($2) and deleted_at IS NULL and country_id = $3"
	var args = []interface{}{time.Now().UTC(), pq.Array(ids), countryId}

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		l.Error("failed to delete individuals", zap.Error(err))
		return err
	}

	return nil
}

func NewIndividualRepo(db *sqlx.DB) IndividualRepo {
	return &individualRepo{db: db}
}
