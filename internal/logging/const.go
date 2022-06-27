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
			fields = append(fields, zap.String("request_id", rid))
		}

		user := utils.GetRequestUser(ctx)
		if user != nil {
			if user.Subject != "" {
				fields = append(fields, zap.String("user_subject", user.Subject))
			}
			if user.Email != "" {
				fields = append(fields, zap.String("user_email", user.Email))
			}
		}

	}
	return l.With(fields...)
}
