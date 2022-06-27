package db

import (
	"context"
	"github.com/jmoiron/sqlx"
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
	var ret []string
	err := u.db.SelectContext(ctx, &ret, "select code from countries where id in (select country_id id from user_countries where user_id = ?)", userID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (u userCountryRepo) SetCountryIDsForUser(ctx context.Context, userID string, countryIDs []string) error {
	_, err := doInTransaction(ctx, u.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		_, err := tx.ExecContext(ctx, "delete from user_countries where user_id = ?", userID)
		if err != nil {
			return nil, err
		}
		if len(countryIDs) == 0 {
			return nil, nil
		}
		for _, countryID := range countryIDs {
			_, err := tx.ExecContext(ctx, "insert into user_countries (user_id, country_id) values ($1, $2)", userID, countryID)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}
