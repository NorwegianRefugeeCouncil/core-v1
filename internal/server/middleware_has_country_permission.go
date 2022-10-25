package server

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func hasCountryPermissionMiddleware(permission auth.Permission) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			l := logging.NewLogger(ctx)

			selectedCountryID, err := utils.GetSelectedCountryID(ctx)
			if err != nil {
				l.Error("failed to get selected country id", zap.Error(err))
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			authInterface, err := utils.GetAuthContext(ctx)
			if err != nil {
				l.Error("failed to get auth context", zap.Error(err))
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			if len(selectedCountryID) > 0 {
				if !authInterface.HasCountryLevelPermission(selectedCountryID, permission) {
					l.Error("user does not have permission to access country", zap.String("country_id", selectedCountryID))
					http.Error(w, "forbidden", http.StatusForbidden)
					return
				}
			}

			h.ServeHTTP(w, r)
		})
	}
}
