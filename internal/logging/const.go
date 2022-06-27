package logging

import (
	"context"

	"go.uber.org/zap"
)

const (
	keyRequestID          = "__request_id"
	keyRequestUserSubject = "__request_user_subject"
	keyRequestUserEmail   = "__request_user_email"
)

func NewLogger(ctx context.Context) *zap.Logger {
	l := logger
	var fields []zap.Field
	if ctx != nil {

		// request id
		if rid := ctx.Value(keyRequestID); rid != nil {
			fields = append(fields, zap.String("request_id", rid.(string)))
		}

		// request user subject
		if subject := ctx.Value(keyRequestUserSubject); subject != nil {
			fields = append(fields, zap.String("request_user_subject", subject.(string)))
		}

		// request user email
		if email := ctx.Value(keyRequestUserEmail); email != nil {
			fields = append(fields, zap.String("request_user_email", email.(string)))
		}

	}
	return l.With(fields...)
}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyRequestID, id)
}

func WithRequestUserSubject(ctx context.Context, subject string) context.Context {
	return context.WithValue(ctx, keyRequestUserSubject, subject)
}

func WithRequestUserEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, keyRequestUserEmail, email)
}

func WithRequestUser(ctx context.Context, subject string, email string) context.Context {
	return WithRequestUserSubject(WithRequestUserEmail(ctx, email), subject)
}
