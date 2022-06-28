package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type IndividualRepo interface {
	GetAll(ctx context.Context, options api.GetAllOptions) ([]*api.Individual, error)
	GetByID(ctx context.Context, id string) (*api.Individual, error)
	Put(ctx context.Context, individual *api.Individual, fields []string) (*api.Individual, error)
	PutMany(ctx context.Context, individuals []*api.Individual, fields []string) ([]*api.Individual, error)
	Delete(ctx context.Context, id string) error
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
		normalizedPhoneNumber := utils.NormalizePhoneNumber(options.PhoneNumber)
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

	if options.Countries != nil {
		if len(options.Countries) == 0 {
			return ret, nil
		}
		countryQueries := []string{}
		if i.driverName() == "sqlite" {
			for _, country := range options.Countries {
				countryQueries = append(countryQueries, "countries LIKE "+nextArg("%\""+country+"\"%"))
			}
			whereClauses = append(whereClauses, "("+strings.Join(countryQueries, " OR ")+")")
		} else if i.driverName() == "postgres" {
			for _, country := range options.Countries {
				countryQueries = append(countryQueries, nextArg(country)+" = ANY (countries)")
			}
			whereClauses = append(whereClauses, "("+strings.Join(countryQueries, " OR ")+")")
		}
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
	err := tx.GetContext(ctx, &ret, "SELECT * FROM individuals WHERE id = $1", id)
	if err != nil {
		l.Error("failed to get individual", zap.Error(err))
		return nil, err
	}
	return &ret, nil
}

func (i individualRepo) PutMany(ctx context.Context, individuals []*api.Individual, fields []string) ([]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		var ret = make([]*api.Individual, len(individuals))
		for k, individual := range individuals {
			ind, err := i.putInternal(ctx, tx, individual, fields)
			if err != nil {
				return nil, err
			}
			ret[k] = ind
		}
		return ret, nil
	})
	if err != nil {
		return nil, err
	}
	return ret.([]*api.Individual), nil
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
	l := logging.NewLogger(ctx)
	l.Debug("putting individual")
	fieldMap := map[string]bool{"id": true}
	for _, field := range fields {
		if field == "phone_number" {
			fieldMap["normalized_phone_number"] = true
		}
		fieldMap[field] = true
	}
	fields = make([]string, 0, len(fieldMap))
	for field := range fieldMap {
		fields = append(fields, field)
	}
	sort.Strings(fields)

	isNew := true
	if len(individual.ID) != 0 {
		_, err := i.getByIdInternal(ctx, tx, individual.ID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				l.Error("failed to get individual", zap.Error(err))
				return nil, err
			}
			individual.ID = xid.New().String()
			isNew = true
		} else {
			isNew = false
		}
	} else {
		individual.ID = xid.New().String()
	}

	if isNew {

		var args []interface{}
		statement := "INSERT INTO individuals (" + strings.Join(fields, ",") + ") VALUES ("
		for i, field := range fields {
			if i != 0 {
				statement += ","
			}
			fieldValue, err := individual.GetFieldValue(field)
			if err != nil {
				l.Error("failed to get field value", zap.String("field", field), zap.Error(err))
				return nil, err
			}
			args = append(args, fieldValue)
			statement += "$" + strconv.Itoa(len(args))
		}
		statement = statement + ")"
		_, err := tx.ExecContext(ctx, statement, args...)
		if err != nil {
			l.Error("failed to insert individual", zap.Error(err))
			return nil, err
		}

	} else {

		var args []interface{}
		statement := "UPDATE individuals SET "
		for i, field := range fields {
			if field == "id" {
				continue
			}
			if i != 0 {
				statement = statement + ","
			}
			fieldValue, err := individual.GetFieldValue(field)
			if err != nil {
				l.Error("failed to get field value", zap.String("field", field), zap.Error(err))
				return nil, err
			}
			args = append(args, fieldValue)
			statement = statement + field + "=$" + strconv.Itoa(len(args))
		}
		args = append(args, individual.ID)
		statement = statement + " WHERE id = $" + strconv.Itoa(len(args))
		_, err := tx.ExecContext(ctx, statement, args...)
		if err != nil {
			l.Error("failed to update individual", zap.Error(err))
			return nil, err
		}

	}

	return i.getByIdInternal(ctx, tx, individual.ID)
}

func (i individualRepo) Delete(ctx context.Context, id string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.deleteInternal(ctx, tx, id)
		return err, nil
	})
	return err
}

func (i individualRepo) deleteInternal(ctx context.Context, tx *sqlx.Tx, id string) error {
	l := logging.NewLogger(ctx).With(zap.String("individual_id", id))
	l.Debug("deleting individual")

	const query = "DELETE FROM individuals WHERE id = ?"
	var args = []interface{}{id}

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		l.Error("failed to delete individual", zap.Error(err))
		return err
	}
	return nil
}

func NewIndividualRepo(db *sqlx.DB) IndividualRepo {
	return &individualRepo{db: db}
}
