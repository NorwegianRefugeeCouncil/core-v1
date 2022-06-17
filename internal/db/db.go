package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func doInTransaction(ctx context.Context, db *sqlx.DB, f func(ctx context.Context, tx *sqlx.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	var retError *error
	rollback := true
	defer func() {
		if rollback {
			if err := tx.Rollback(); err != nil {
				retError = &err
			}
		} else {
			if err := tx.Commit(); err != nil {
				*retError = err
			}
		}
	}()
	ret, err := f(ctx, tx)
	if err != nil {
		return nil, err
	}
	rollback = false
	return ret, err
}
