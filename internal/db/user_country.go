package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

type UserCountryRepo interface {
	GetForUserID(ctx context.Context, userID string) ([]string, error)
	SetCountryIDsForUser(ctx context.Context, userID string, countryCodes []string) error
}

type userCountryRepo struct {
	db *sqlx.DB
}

func NewUserCountryRepo(db *sqlx.DB) UserCountryRepo {
	return &userCountryRepo{db: db}
}

func (u userCountryRepo) GetForUserID(ctx context.Context, userID string) ([]string, error) {
	l := logging.NewLogger(ctx).With(zap.String("user_id", userID))
	l.Debug("getting user countries")

	const query = "select code from countries where id in (select country_id id from user_countries where user_id = ?)"
	var args = []interface{}{userID}

	var ret []string
	if err := u.db.SelectContext(ctx, &ret, query, args...); err != nil {
		l.Error("failed to get countries for user", zap.Error(err))
		return nil, err
	}
	return ret, nil
}

func (u userCountryRepo) SetCountryIDsForUser(ctx context.Context, userID string, countryIDs []string) error {
	l := logging.NewLogger(ctx).With(zap.String("user_id", userID))
	l.Debug("setting user countries")
	_, err := doInTransaction(ctx, u.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {

		const deleteQuery = "delete from user_countries where user_id = ?"
		var deleteArgs = []interface{}{userID}

		if _, err := tx.ExecContext(ctx, deleteQuery, deleteArgs...); err != nil {
			l.Error("failed to delete user countries", zap.Error(err))
			return nil, err
		}
		if len(countryIDs) == 0 {
			return nil, nil
		}
		for _, countryID := range countryIDs {

			const insertQuery = "insert into user_countries (user_id, country_id, permission) values ($1, $2, $3)"
			var insertArgs = []interface{}{
				userID,
				countryID,
				"",
			}

			if _, err := tx.ExecContext(ctx, insertQuery, insertArgs...); err != nil {
				l.Error("failed to insert user country", zap.Error(err))
				return nil, err
			}
		}
		return nil, nil
	})
	if err != nil {
		l.Error("failed to set user countries", zap.Error(err))
		return err
	}
	return nil
}
