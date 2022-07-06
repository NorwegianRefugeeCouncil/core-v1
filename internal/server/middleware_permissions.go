package server

import (
	"github.com/nrc-no/notcore/internal/clients/zanzibar"
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func permissionMiddleware(permissionRepo db.PermissionRepo, client *zanzibar.ZanzibarClient) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			user := utils.GetRequestUser(ctx)
			if user == nil {
				l.Error("user not found in context")
				http.Error(w, "Invalid authorization header", http.StatusBadRequest)
				return
			}

			permission, err := permissionRepo.GetPermissionsForUser(r.Context(), user.ID)
			if err != nil {
				l.Error("couldn't get permissions for user: ", zap.Error(err))
				http.Error(w, "couldn't get permissions: "+err.Error(), http.StatusInternalServerError)
				return
			}

			isGlobalAdmin, err := client.CheckIsGlobalAdmin(ctx)
			permission.IsGlobalAdmin = isGlobalAdmin
			ctx = utils.WithUserPermissions(ctx, *permission)
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}
