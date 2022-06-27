package utils

import (
	"context"

	"github.com/nrc-no/notcore/internal/api"
)

const (
	keyRequestID              = "__request_id"
	keyRequestUserSubject     = "__request_user_subject"
	keyRequestUserEmail       = "__request_user_email"
	keyRequestUser            = "__request_user"
	keyRequestUserPermissions = "__request_user_permissions"
)

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyRequestID, id)
}

func GetRequestID(ctx context.Context) string {
	if rid := ctx.Value(keyRequestID); rid != nil {
		return rid.(string)
	}
	return ""
}

func WithUser(ctx context.Context, user api.User) context.Context {
	ctx = context.WithValue(ctx, keyRequestUser, user)
	return ctx
}

func GetRequestUser(ctx context.Context) *api.User {
	if ctx == nil {
		return nil
	}
	if user := ctx.Value(keyRequestUser); user != nil {
		if u, ok := user.(api.User); ok {
			return &u
		}
	}
	return nil
}

func WithUserPermissions(ctx context.Context, userPermissions api.UserPermissions) context.Context {
	ctx = context.WithValue(ctx, keyRequestUserPermissions, userPermissions)
	return ctx
}

func GetRequestUserPermissions(ctx context.Context) api.UserPermissions {
	if userPermissions := ctx.Value(keyRequestUserPermissions); userPermissions != nil {
		if up, ok := userPermissions.(api.UserPermissions); ok {
			return up
		}
	}
	return api.UserPermissions{}
}
