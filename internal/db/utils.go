package db

import (
	"context"
	"fmt"
	"time"

	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func batch[T any](batchSize int, all []T, fn func([]T) (stop bool, err error)) error {
	if batchSize <= 0 {
		return fmt.Errorf("batch size must be greater than 0")
	}
	count := len(all)
	for b := 0; b < count; b += batchSize {
		stop, err := fn(all[b:utils.Min(b+batchSize, count)])
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

func logDuration(ctx context.Context, name string, fields ...zap.Field) func() {
	start := time.Now()
	logger := logging.NewLogger(ctx)
	return func() {
		fields = append([]zap.Field{
			zap.String("operation", name),
			zap.Duration("duration", time.Since(start)),
		}, fields...)
		logger.Debug("duration audit", fields...)
	}
}
