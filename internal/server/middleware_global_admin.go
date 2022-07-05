package server

import (
	"github.com/nrc-no/notcore/internal/clients/zanzibar"
	"github.com/nrc-no/notcore/internal/utils"
	"net/http"
	"sync"

	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

func firstUserGlobalAdmin(permissionRepo db.PermissionRepo, client zanzibar.Client) func(h http.Handler) http.Handler {

	foundGlobalAdmin := atomic.NewBool(false)
	lock := &sync.Mutex{}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := logging.NewLogger(r.Context())
			if !foundGlobalAdmin.Load() {
				l.Info("checking for global admin")
				if lock.TryLock() {
					defer lock.Unlock()

					hasAny, err := client.Permissions.CheckGlobalAdminExists(r.Context())
					if err != nil {
						l.Error("zanzibar: couldn't check if there is a global admin: ", zap.Error(err))
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					if err != nil {
						l.Error("couldn't check if there is a global admin: ", zap.Error(err))
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					if !hasAny {
						l.Info("no global admin found, assigning the first user as global admin")

						// create global admin in zanzibar
						_, err := client.Relation.AddGlobalAdmin(r.Context())
						if err != nil {
							l.Error("zanzibar: couldn't save global admin relation for user: ", zap.Error(err))
							return
						}

						//// adds admin permission to user
						perms := utils.GetRequestUserPermissions(r.Context())
						perms.IsGlobalAdmin = true
						//if err := permissionRepo.SavePermissionsForUser(r.Context(), &perms); err != nil {
						//	http.Error(w, err.Error(), http.StatusInternalServerError)
						//	return
						//}
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
