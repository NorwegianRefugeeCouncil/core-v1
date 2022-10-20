package logging

import (
	"context"

	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func NewLogger(ctx context.Context) *zap.Logger {
	l := logger
	var fields []zap.Field
	if ctx != nil {

		// request id
		rid := utils.GetRequestID(ctx)
		if rid != "" {
			fields = append(fields, zap.String("rid", rid))
		}

		session, ok := utils.GetSession(ctx)
		if ok {
			fields = append(fields, zap.String("user", session.GetUserID()))
		}

	}
	return l.With(fields...)
}
