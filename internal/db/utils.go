package db

import (
	"context"
	"fmt"
	"github.com/nrc-no/notcore/internal/constants"
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

func GetDeduplicationOptionNames(deduplicationTypes []string) ([]DeduplicationOptionName, []string, error) {
	optionNames := make([]DeduplicationOptionName, 0)
	fileColumns := make([]string, 0)
	for _, d := range deduplicationTypes {
		dt, ok := ParseDeduplicationOptionName(d)
		optionNames = append(optionNames, dt)
		if ok {
			for _, vc := range DeduplicationOptions[dt].Value.Columns {
				fileColumns = append(fileColumns, constants.IndividualDBToFileMap[vc])
			}
		} else {
			return nil, nil, fmt.Errorf("invalid deduplication type: %s", d)
		}
	}
	return optionNames, fileColumns, nil
}
