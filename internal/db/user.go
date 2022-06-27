package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/rs/xid"
	"strconv"
)

type UserRepo interface {
	GetByID(ctx context.Context, userID string) (*api.User, error)
	Put(ctx context.Context, user *api.User) (*api.User, error)
	GetAll(ctx context.Context, options api.GetAllUsersOptions) ([]*api.User, error)
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

type userRepo struct {
	db *sqlx.DB
}

func (r userRepo) GetByID(ctx context.Context, userID string) (*api.User, error) {
	ret := api.User{}
	err := r.db.GetContext(ctx, &ret, "select * from users where id = $1", userID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (r userRepo) Put(ctx context.Context, user *api.User) (*api.User, error) {
	ret, err := doInTransaction(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return r.put(ctx, tx, user)
	})
	if err != nil {
		return nil, err
	}
	return ret.(*api.User), nil
}

func (r userRepo) put(ctx context.Context, tx *sqlx.Tx, user *api.User) (*api.User, error) {
	found, err := r.findSubject(ctx, tx, user.Subject)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		user.ID = xid.New().String()
		_, err = tx.ExecContext(ctx, "insert into users (id, subject, email) values($1, $2, $3)",
			user.ID,
			user.Subject,
			user.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return found, nil
}

func (r userRepo) findSubject(ctx context.Context, tx *sqlx.Tx, subject string) (*api.User, error) {
	var ret api.User
	err := tx.GetContext(ctx, &ret, `SELECT * FROM users WHERE subject = $1`, subject)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (r userRepo) GetAll(ctx context.Context, options api.GetAllUsersOptions) ([]*api.User, error) {
	ret, err := doInTransaction(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return r.getAll(ctx, tx, options)
	})
	if err != nil {
		return nil, err
	}
	return ret.([]*api.User), nil
}

func (r userRepo) getAll(ctx context.Context, tx *sqlx.Tx, options api.GetAllUsersOptions) ([]*api.User, error) {

	qry := "select * from users"
	var whereClauses []string
	var args []interface{}

	if len(options.CountryIDs) != 0 {
		countryQry := "id in (select user_id from user_countries where country_id in ("
		for i, countryID := range options.CountryIDs {
			if i == 0 {
				countryQry += ","
			}
			args = append(args, countryID)
			countryQry += fmt.Sprintf("$%d", len(args))
		}
		countryQry += "))"
		whereClauses = append(whereClauses, countryQry)
	}

	if len(options.Email) != 0 {
		args = append(args, "%"+options.Email+"%")
		emailQry := "email like $" + strconv.Itoa(len(args))
		whereClauses = append(whereClauses, emailQry)
	}

	if len(whereClauses) != 0 {
		for i, whereClause := range whereClauses {
			if i != 0 {
				qry += " AND ("
			} else {
				qry += " ("
			}
			qry += whereClause
			qry += ")"
		}
	}

	if options.Take != 0 {
		qry += fmt.Sprintf(" LIMIT %d", options.Take)
	}

	if options.Skip != 0 {
		qry += fmt.Sprintf(fmt.Sprintf(" OFFSET %d", options.Skip))
	}

	var ret []*api.User
	err := tx.SelectContext(ctx, &ret, qry, args...)
	if err != nil {
		return nil, err
	}
	return ret, nil

}
