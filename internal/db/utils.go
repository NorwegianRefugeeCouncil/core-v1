package db

import (
	"fmt"

	"github.com/nrc-no/notcore/internal/utils"
)

func batch[T any](batchSize int, all []T, fn func([]T) error) error {
	if batchSize <= 0 {
		return fmt.Errorf("batch size must be greater than 0")
	}

	count := len(all)

	if count == 0 {
		return fn(all)
	}

	for b := 0; b < count; b += batchSize {
		if err := fn(all[b:utils.Min(b+batchSize, count)]); err != nil {
			return err
		}
	}
	return nil
}
