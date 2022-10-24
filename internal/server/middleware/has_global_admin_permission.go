package middleware

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HasGlobalAdminPermission() func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			l := logging.NewLogger(ctx)

			authInterface, err := utils.GetAuthContext(ctx)
			if err != nil {
				l.Error("failed to get auth context", zap.Error(err))
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			if !authInterface.IsGlobalAdmin() {
				l.Error("user is not global admin")
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
