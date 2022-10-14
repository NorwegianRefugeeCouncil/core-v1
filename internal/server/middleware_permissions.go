package server

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func permissionMiddleware(permissionRepo db.PermissionRepo) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			l.Debug("permission middleware")

			user := utils.GetRequestUser(ctx)
			if user == nil {
				l.Error("user not found in context")
				http.Error(w, "Invalid authorization header", http.StatusBadRequest)
				return
			}

			permission, err := permissionRepo.GetExplicitPermissionsForUser(r.Context(), user.ID)
			if err != nil {
				l.Error("couldn't get permissions for user: ", zap.Error(err))
				http.Error(w, "couldn't get permissions: "+err.Error(), http.StatusInternalServerError)
				return
			}

			allCountries, err := utils.GetCountries(ctx)
			if err != nil {
				l.Error("failed to get all countries", zap.Error(err))
				http.Error(w, "couldn't get countries: "+err.Error(), http.StatusInternalServerError)
				return
			}

			authIntf := auth.New(*permission, allCountries)
			r = r.WithContext(utils.WithAuthContext(ctx, authIntf))
			h.ServeHTTP(w, r)
		})
	}
}
