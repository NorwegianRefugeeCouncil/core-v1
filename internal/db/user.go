package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/rs/xid"
	"go.uber.org/zap"
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
	l := logging.NewLogger(ctx).With(zap.String("user_id", userID))
	l.Debug("getting user by id")

	const query = "select * from users where id = $1"
	var args = []interface{}{userID}

	ret := api.User{}
	if err := r.db.GetContext(ctx, &ret, query, args...); err != nil {
		l.Error("failed to get user by id", zap.Error(err))
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
	l := logging.NewLogger(ctx)

	l.Debug("putting user")
	found, err := r.findSubject(ctx, tx, user.Subject)
	if err == nil {
		return found, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		l.Error("failed to find user", zap.Error(err))
		return nil, err
	}

	user.ID = xid.New().String()

	const query = "insert into users (id, subject, email) values($1, $2, $3)"
	var args = []interface{}{user.ID, user.Subject, user.Email}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		l.Error("failed to insert user", zap.Error(err))
		return nil, err
	}

	return user, nil

}

func (r userRepo) findSubject(ctx context.Context, tx *sqlx.Tx, subject string) (*api.User, error) {
	l := logging.NewLogger(ctx).With(zap.String("subject", subject))
	l.Debug("finding user by subject")

	const query = `SELECT * FROM users WHERE subject = $1`
	var args = []interface{}{subject}

	var ret api.User
	if err := tx.GetContext(ctx, &ret, query, args...); err != nil {
		l.Error("failed to find user", zap.Error(err))
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

	l := logging.NewLogger(ctx)
	l.Debug("getting all users", zap.Any("options", options))

	qry := "select * from users"
	var whereClauses []string
	var args []interface{}

	args, whereClauses = r.withCountryIdsClause(options, args, whereClauses)
	args, whereClauses = r.withEmailClause(options, args, whereClauses)

	qry = r.withWhereClauses(whereClauses, qry)
	qry = r.withQueryTake(options, qry)
	qry = r.withQueryOffset(options, qry)

	var ret []*api.User
	err := tx.SelectContext(ctx, &ret, qry, args...)
	if err != nil {
		l.Error("failed to get all users", zap.Error(err))
		return nil, err
	}
	return ret, nil

}

func (r userRepo) withWhereClauses(whereClauses []string, qry string) string {
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
	return qry
}

func (r userRepo) withCountryIdsClause(options api.GetAllUsersOptions, args []interface{}, whereClauses []string) ([]interface{}, []string) {
	if len(options.CountryIDs) != 0 {
		qry := "id in (select user_id from user_countries where country_id in ("
		for i, countryID := range options.CountryIDs {
			if i == 0 {
				qry += ","
			}
			args = append(args, countryID)
			qry += fmt.Sprintf("$%d", len(args))
		}
		qry += "))"
		whereClauses = append(whereClauses, qry)
	}
	return args, whereClauses
}

func (r userRepo) withEmailClause(options api.GetAllUsersOptions, args []interface{}, whereClauses []string) ([]interface{}, []string) {
	if len(options.Email) != 0 {
		args = append(args, "%"+options.Email+"%")
		emailQry := "email like $" + strconv.Itoa(len(args))
		whereClauses = append(whereClauses, emailQry)
	}
	return args, whereClauses
}

func (r userRepo) withQueryOffset(options api.GetAllUsersOptions, qry string) string {
	if options.Skip != 0 {
		qry += fmt.Sprintf(fmt.Sprintf(" OFFSET %d", options.Skip))
	}
	return qry
}

func (r userRepo) withQueryTake(options api.GetAllUsersOptions, qry string) string {
	if options.Take != 0 {
		qry += fmt.Sprintf(" LIMIT %d", options.Take)
	}
	return qry
}
