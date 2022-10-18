package utils

import (
	"context"
	"fmt"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
)

const (
	keyRequestID              = "__request_id"
	keyRequestUserSubject     = "__request_user_subject"
	keyRequestUserEmail       = "__request_user_email"
	keyRequestUser            = "__request_user"
	keyRequestUserPermissions = "__request_user_permissions"
	keyAuthContext            = "__request_auth"
	keyCountries              = "__request_countries"
	keySelectedCountryID      = "__request_selected_country"
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

func WithAuthContext(ctx context.Context, authCtx auth.Interface) context.Context {
	ctx = context.WithValue(ctx, keyAuthContext, authCtx)
	return ctx
}

func GetAuthContext(ctx context.Context) (auth.Interface, error) {
	if authCtxIntf := ctx.Value(keyAuthContext); authCtxIntf != nil {
		if authCtx, ok := authCtxIntf.(auth.Interface); ok {
			return authCtx, nil
		}
		return nil, fmt.Errorf("failed to get auth context: wrong value type")
	}
	return nil, fmt.Errorf("failed to get auth context: value not present")
}

func WithCountries(ctx context.Context, countries []*api.Country) context.Context {
	ctx = context.WithValue(ctx, keyCountries, countries)
	return ctx
}

func GetCountries(ctx context.Context) ([]*api.Country, error) {
	if countriesIntf := ctx.Value(keyCountries); countriesIntf != nil {
		if countries, ok := countriesIntf.([]*api.Country); ok {
			return countries, nil
		}
		return nil, fmt.Errorf("failed to get countries: wrong value type")
	}
	return nil, fmt.Errorf("failed to get countries: value not present")
}

func WithSelectedCountryID(ctx context.Context, selectedCountryID string) context.Context {
	ctx = context.WithValue(ctx, keySelectedCountryID, selectedCountryID)
	return ctx
}

func GetSelectedCountryID(ctx context.Context) (string, error) {
	if selectedCountryIDIntf := ctx.Value(keySelectedCountryID); selectedCountryIDIntf != nil {
		if countries, ok := selectedCountryIDIntf.(string); ok {
			return countries, nil
		}
		return "", fmt.Errorf("failed to get selected country id: wrong value type")
	}
	return "", nil
}
