package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/rs/xid"
)

type IndividualRepo interface {
	GetAll(ctx context.Context, options api.GetAllOptions) ([]*api.Individual, error)
	GetByID(ctx context.Context, id string) (*api.Individual, error)
	Put(ctx context.Context, individual *api.Individual) (*api.Individual, error)
	PutMany(ctx context.Context, individuals []*api.Individual) ([]*api.Individual, error)
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
			whereClause := fmt.Sprintf("rowid IN (SELECT rowid FROM individuals_fts WHERE full_name MATCH %s ORDER BY rank)", nextArg(options.FullName))
			whereClauses = append(whereClauses, whereClause)
		} else if i.driverName() == "postgres" {
			whereClause := "full_name ILIKE " + nextArg("%"+options.FullName+"%")
			whereClauses = append(whereClauses, whereClause)
		}
	}

	if len(options.Address) != 0 {
		if i.driverName() == "sqlite" {
			whereClause := fmt.Sprintf("rowid IN (SELECT rowid FROM individuals_fts WHERE address MATCH %s ORDER BY rank)", nextArg(options.Address))
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

	if len(whereClauses) != 0 {
		query = query + " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query = query + " ORDER BY ID"
	if options.Take != 0 {
		query += fmt.Sprintf(" LIMIT %d", options.Take)
	}
	if options.Skip != 0 {
		query += fmt.Sprintf(fmt.Sprintf(" OFFSET %d", options.Skip))
	}
	err := tx.SelectContext(ctx, &ret, query, args...)
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
	var ret = api.Individual{}
	err := tx.GetContext(ctx, &ret, "SELECT * FROM individuals WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (i individualRepo) PutMany(ctx context.Context, individuals []*api.Individual) ([]*api.Individual, error) {

	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		var ret = make([]*api.Individual, len(individuals))
		for k, individual := range individuals {
			ind, err := i.putInternal(ctx, tx, individual)
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

func (i individualRepo) Put(ctx context.Context, individual *api.Individual) (*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.putInternal(ctx, tx, individual)
	})
	if err != nil {
		return nil, err
	}
	return ret.(*api.Individual), nil
}

func (i individualRepo) putInternal(ctx context.Context, tx *sqlx.Tx, individual *api.Individual) (*api.Individual, error) {

	isNew := true
	if len(individual.ID) != 0 {
		_, err := i.getByIdInternal(ctx, tx, individual.ID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
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
		_, err := tx.ExecContext(ctx, `
INSERT INTO individuals (id, full_name, email, phone_number, normalized_phone_number, address, birth_date, gender) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`,
			individual.ID,
			individual.FullName,
			individual.Email,
			individual.PhoneNumber,
			individual.NormalizedPhoneNumber,
			individual.Address,
			individual.BirthDate,
			individual.Gender,
		)
		if err != nil {
			return nil, err
		}

	} else {
		_, err := tx.ExecContext(ctx, `
UPDATE individuals SET 
full_name = $1,
email = $2, 
phone_number = $3, 
normalized_phone_number = $4,
address = $5,
birth_date = $6,
gender = $7
WHERE id = $8`,
			individual.FullName,
			individual.Email,
			individual.PhoneNumber,
			individual.NormalizedPhoneNumber,
			individual.Address,
			individual.BirthDate,
			individual.Gender,
			individual.ID,
		)
		if err != nil {
			return nil, err
		}
	}

	return individual, nil
}

func (i individualRepo) Delete(ctx context.Context, id string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.deleteInternal(ctx, tx, id)
		return err, nil
	})
	return err
}

func (i individualRepo) deleteInternal(ctx context.Context, tx *sqlx.Tx, id string) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM individuals WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func NewIndividualRepo(db *sqlx.DB) IndividualRepo {
	return &individualRepo{db: db}
}
