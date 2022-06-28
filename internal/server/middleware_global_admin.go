package server

import (
	"net/http"
	"sync"

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
			l := logging.NewLogger(r.Context())
			if !foundGlobalAdmin.Load() {
				l.Info("checking for global admin")
				if lock.TryLock() {
					defer lock.Unlock()
					hasAny, err := permissionRepo.HasAnyGlobalAdmin(r.Context())
					if err != nil {
						l.Error("couldn't check if there is a global admin: ", zap.Error(err))
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					if !hasAny {
						l.Info("no global admin found, assigning the first user as global admin")
						perms := utils.GetRequestUserPermissions(r.Context())
						perms.IsGlobalAdmin = true
						if err := permissionRepo.SavePermissionsForUser(r.Context(), &perms); err != nil {
							l.Error("couldn't save permissions for user: ", zap.Error(err))
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						r = r.WithContext(utils.WithUserPermissions(r.Context(), perms))
						foundGlobalAdmin.Store(true)
					} else {
						l.Info("global admin found")
						foundGlobalAdmin.Store(true)
					}
				}
			}
			h.ServeHTTP(w, r)
		})
	}
}
