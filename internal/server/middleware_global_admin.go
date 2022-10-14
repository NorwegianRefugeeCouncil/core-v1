package server

import (
	"net/http"
	"sync"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

func firstUserGlobalAdmin(permissionRepo db.PermissionRepo) func(h http.Handler) http.Handler {

	foundGlobalAdmin := atomic.NewBool(false)
	lock := &sync.Mutex{}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			l := logging.NewLogger(ctx)

			if foundGlobalAdmin.Load() {
				h.ServeHTTP(w, r)
				return
			}

			l.Info("checking for global admin")
			if lock.TryLock() {
				defer lock.Unlock()

				isEmpty, err := permissionRepo.IsEmpty(ctx)
				if err != nil {
					l.Error("couldn't check if there already is a user permission: ", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				hasAny := !isEmpty

				if !hasAny {
					l.Info("no global admin found, assigning the first user as global admin")

					requestUser := utils.GetRequestUser(ctx)
					perms := api.UserPermissions{
						UserID:        requestUser.ID,
						IsGlobalAdmin: true,
					}

					if err := permissionRepo.SaveExplicitPermissionsForUser(ctx, &perms); err != nil {
						l.Error("couldn't save permissions for user: ", zap.Error(err))
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					foundGlobalAdmin.Store(true)
				} else {
					l.Info("global admin found")
					foundGlobalAdmin.Store(true)
				}

			}

			h.ServeHTTP(w, r)
		})
	}
}
