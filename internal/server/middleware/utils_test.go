package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils"
)

func nextHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func configureDummyContextMiddleware(
	allowedCountryIDs containers.Set[string],
	allCountryIDs containers.Set[string],
	isGlobalAdmin bool,
	selectedCountryId string,
) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authIntf := auth.New(
				allowedCountryIDs,
				allCountryIDs,
				isGlobalAdmin,
			)
			r = r.WithContext(
				utils.WithAuthContext(
					utils.WithSelectedCountryID(r.Context(), selectedCountryId),
					authIntf,
				),
			)
			next.ServeHTTP(w, r)
		})
	}
}
