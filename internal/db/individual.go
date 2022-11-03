package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
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
	SoftDelete(ctx context.Context, id string) error
	SoftDeleteMany(ctx context.Context, ids []string) error
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
		return i.batchedGetAllInternal(ctx, tx, options)
	})
	if err != nil {

		return nil, err
	}
	return ret.([]*api.Individual), nil
}

func (i individualRepo) batchedGetAllInternal(ctx context.Context, tx *sqlx.Tx, options api.GetAllOptions) ([]*api.Individual, error) {
	if len(options.IDs) > 0 {
		var ret []*api.Individual
		err := batch(maxParams/len(options.IDs), options.IDs, func(idsInBatch []string) error {
			optionsForBatch := options
			optionsForBatch.IDs = idsInBatch
			individualsInBatch, err := i.unbatchedGetAllInternal(ctx, tx, optionsForBatch)
			if err != nil {
				return err
			}
			ret = append(ret, individualsInBatch...)
			return nil
		})
		if err != nil {
			return nil, err
		}
		return ret, nil
	}
	return i.unbatchedGetAllInternal(ctx, tx, options)
}

func (i individualRepo) unbatchedGetAllInternal(ctx context.Context, tx *sqlx.Tx, options api.GetAllOptions) ([]*api.Individual, error) {
	l := logging.NewLogger(ctx)
	l.Debug("getting list individuals", zap.Any("options", options))
	var ret []*api.Individual
	sql, args := newGetAllIndividualsSQLQuery(i.driverName(), options).build()
	err := tx.SelectContext(ctx, &ret, sql, args...)
	if err != nil {
		l.Error("failed to list individuals", zap.Error(err))
	}
	return ret, nil
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
			if field == "id" || field == "country_id" {
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

func (i individualRepo) SoftDelete(ctx context.Context, id string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.softDeleteManyInternal(ctx, tx, []string{id})
		return nil, err
	})
	return err
}

func (i individualRepo) SoftDeleteMany(ctx context.Context, ids []string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.softDeleteManyInternal(ctx, tx, ids)
		return nil, err
	})
	return err
}

func (i individualRepo) softDeleteManyInternal(ctx context.Context, tx *sqlx.Tx, ids []string) error {
	idSet := containers.NewStringSet(ids...)
	ids = idSet.Items()

	l := logging.NewLogger(ctx).With(zap.Strings("individual_ids", ids))
	l.Debug("deleting individuals")

	if err := batch(maxParams/len(ids), ids, func(idsInBatch []string) error {
		var query = "UPDATE individuals SET deleted_at = $1 WHERE id IN ("
		var args = []interface{}{time.Now().UTC().Format(time.RFC3339)}
		for i, id := range idsInBatch {
			if i != 0 {
				query += ","
			}
			query += fmt.Sprintf("$%d", i+2)
			args = append(args, id)
		}
		query += ") and deleted_at IS NULL"

		result, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			l.Error("failed to delete individuals", zap.Error(err))
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			l.Error("failed to get rows affected", zap.Error(err))
			return err
		} else if rowsAffected != int64(len(idsInBatch)) {
			l.Error("failed to delete all individuals", zap.Int64("rows_affected", rowsAffected))
			return fmt.Errorf("failed to delete all individuals")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func NewIndividualRepo(db *sqlx.DB) IndividualRepo {
	return &individualRepo{db: db}
}
