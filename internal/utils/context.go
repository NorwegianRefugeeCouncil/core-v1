package utils

import (
	"context"
	"fmt"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
)

type key uint8

const (
	keyRequestID key = iota
	keyRequestSession
	keyAuthContext
	keyCountries
	keySelectedCountryID
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

func WithSession(ctx context.Context, session auth.Session) context.Context {
	ctx = context.WithValue(ctx, keyRequestSession, session)
	return ctx
}

func GetSession(ctx context.Context) (auth.Session, bool) {
	if ctx == nil {
		return nil, false
	}
	if session := ctx.Value(keyRequestSession); session != nil {
		if s, ok := session.(auth.Session); ok {
			return s, true
		}
	}
	return nil, false
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
