package db

import (
	"context"
	"time"

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
	retError := new(error)
	rollback := true
	defer func() {
		if rollback {
			rollbackStart := time.Now()
			l.Warn("rolling back transaction")
			if err := tx.Rollback(); err != nil {
				l.Error("failed to rollback transaction",
					zap.Error(err),
					zap.Duration("rollback_duration", time.Since(rollbackStart)),
				)
				retError = &err
			} else {
				l.Debug("rolled back transaction",
					zap.Duration("rollback_duration", time.Since(rollbackStart)),
				)
			}
		} else {
			commitStart := time.Now()
			l.Debug("committing transaction")
			if err := tx.Commit(); err != nil {
				l.Error("failed to commit transaction",
					zap.Error(err),
					zap.Duration("commit_duration", time.Since(commitStart)),
				)
				*retError = err
			} else {
				l.Debug("committed transaction",
					zap.Duration("commit_duration", time.Since(commitStart)),
				)
			}
		}
	}()
	transactionStart := time.Now()
	ret, err := f(ctx, tx)
	if err != nil {
		l.Error("failed to execute transaction",
			zap.Error(err),
			zap.Duration("transaction_duration", time.Since(transactionStart)),
		)
		return nil, err
	} else {
		l.Debug("executed transaction",
			zap.Duration("transaction_duration", time.Since(transactionStart)),
		)
	}
	rollback = false
	return ret, err
}
