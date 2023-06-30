package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

type HealthzRepo interface {
	Check(ctx context.Context) (error)
}

type healthzRepo struct {
	db *sqlx.DB
}

func NewHealthzRepo(db *sqlx.DB) HealthzRepo {
	return &healthzRepo{db: db}
}

func (h healthzRepo) logger(ctx context.Context) *zap.Logger {
	return logging.NewLogger(ctx)
}

func (h healthzRepo) Check(ctx context.Context) error {
	l := h.logger(ctx)

	if err := h.db.PingContext(ctx); err != nil {
		l.Error("failed to ping database", zap.Error(err))
		return err
	}

	return nil
}