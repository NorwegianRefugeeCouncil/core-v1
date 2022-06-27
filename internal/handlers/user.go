package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleUser(
	templates map[string]*template.Template,
	countryRepo db.CountryRepo,
	userRepo db.UserRepo,
	permissionRepo db.PermissionRepo,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const (
			templateName         = "user.gohtml"
			pathParamUserId      = "user_id"
			viewParamUser        = "User"
			viewParamPermissions = "Permissions"
			viewParamCountries   = "Countries"
		)

		var (
			ctx         = r.Context()
			userID      = mux.Vars(r)[pathParamUserId]
			user        *api.User
			permissions *api.UserPermissions
			countries   []*api.Country
			l           = logging.NewLogger(ctx)
		)

		render := func() {
			renderView(templates, templateName, w, r, map[string]interface{}{
				viewParamUser:        user,
				viewParamPermissions: permissions,
				viewParamCountries:   countries,
			})
		}

		group, gtx := errgroup.WithContext(ctx)
		group.Go(func() error {
			var err error
			if user, err = userRepo.GetByID(gtx, userID); err != nil {
				l.Error("couldn't get user: ", zap.Error(err))
				return err
			}
			return nil
		})
		group.Go(func() error {
			var err error
			if permissions, err = permissionRepo.GetPermissionsForUser(gtx, userID); err != nil {
				l.Error("couldn't get permissions for user: ", zap.Error(err))
				return err
			}
			return nil
		})
		group.Go(func() error {
			var err error
			if countries, err = countryRepo.GetAll(gtx); err != nil {
				l.Error("couldn't get countries: ", zap.Error(err))
				return err
			}
			return nil
		})
		if err := group.Wait(); err != nil {
			http.Error(w, "couldn't get user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Method == http.MethodGet {
			render()
			return
		}

		if err := r.ParseForm(); err != nil {
			l.Error("couldn't parse form: ", zap.Error(err))
			http.Error(w, "couldn't parse form: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var countryPermissions = api.CountryPermissions{}

		for key, values := range r.Form {
			if strings.HasPrefix(key, "permissions[") {
				countryID := strings.TrimPrefix(key, "permissions[")
				countryID = strings.TrimSuffix(countryID, "]")
				if _, ok := countryPermissions[countryID]; !ok {
					countryPermissions[countryID] = api.CountryPermission{
						CountryID: countryID,
					}
				}
				var permission = countryPermissions[countryID]
				for _, value := range values {
					if value == "read" {
						permission.Read = true
					}
					if value == "write" {
						permission.Write = true
					}
					if value == "admin" {
						permission.Admin = true
					}
				}
				countryPermissions[countryID] = permission
			}
		}

		permissions = &api.UserPermissions{
			UserID:             userID,
			CountryPermissions: countryPermissions,
		}
		if err := permissionRepo.SavePermissionsForUser(ctx, permissions); err != nil {
			l.Error("couldn't save permissions for user: ", zap.Error(err))
			http.Error(w, "couldn't save permissions for user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		render()

	})
}
