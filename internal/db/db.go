package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func doInTransaction(ctx context.Context, db *sqlx.DB, f func(ctx context.Context, tx *sqlx.Tx) (interface{}, error)) (interface{}, error) {
	l := logging.NewLogger(ctx)
	l.Debug("starting transaction")
	tx, err := db.Beginx()
	if err != nil {
		l.Error("failed to begin transaction", zap.Error(err))
		return nil, err
	}
	var retError *error
	rollback := true
	defer func() {
		if rollback {
			l.Warn("rolling back transaction")
			if err := tx.Rollback(); err != nil {
				l.Error("failed to rollback transaction", zap.Error(err))
				retError = &err
			}
		} else {
			l.Debug("committing transaction")
			if err := tx.Commit(); err != nil {
				l.Error("failed to commit transaction", zap.Error(err))
				*retError = err
			}
		}
	}()
	ret, err := f(ctx, tx)
	if err != nil {
		l.Error("failed to execute transaction", zap.Error(err))
		return nil, err
	}
	rollback = false
	return ret, err
}
