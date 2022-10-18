package handlers

import (
	"errors"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleUser(
	templates map[string]*template.Template,
	userRepo db.UserRepo,
	permissionRepo db.PermissionRepo,
) http.Handler {

	const (
		templateName         = "user.gohtml"
		pathParamUserId      = "user_id"
		viewParamUser        = "User"
		viewParamPermissions = "Permissions"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx         = r.Context()
			l           = logging.NewLogger(ctx)
			userID      = mux.Vars(r)[pathParamUserId]
			user        *api.User
			permissions *api.UserPermissions
			authIntf    auth.Interface
			err         error
		)

		render := func() {
			renderView(templates, templateName, w, r, viewParams{
				viewParamUser:        user,
				viewParamPermissions: permissions,
			})
		}

		errGrp, gtx := errgroup.WithContext(ctx)
		errGrp.Go(func() error {
			var err error
			if user, err = userRepo.GetByID(gtx, userID); err != nil {
				l.Error("couldn't get user: ", zap.Error(err))
				return err
			}
			return nil
		})
		errGrp.Go(func() error {
			var err error
			if permissions, err = permissionRepo.GetExplicitPermissionsForUser(gtx, userID); err != nil {
				l.Error("couldn't get permissions for user: ", zap.Error(err))
				return err
			}
			return nil
		})
		if err := errGrp.Wait(); err != nil {
			http.Error(w, "couldn't get user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		authIntf, err = utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		countryIdsWhereAdmin := authIntf.GetCountryIDsWithPermission("admin")

		if !authIntf.IsGlobalAdmin() && len(countryIdsWhereAdmin) == 0 {
			http.Error(w, "you don't have sufficient permissions", http.StatusForbidden)
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

		if err := parseUserForm(r, authIntf, permissions); err != nil {
			http.Error(w, "you don't have sufficient permissions", http.StatusForbidden)
			return
		}

		if err := permissionRepo.SaveExplicitPermissionsForUser(ctx, permissions); err != nil {
			l.Error("couldn't save permissions for user: ", zap.Error(err))
			http.Error(w, "couldn't save permissions for user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		render()

	})
}

func filterCountriesWhereAdmin(permHelper auth.Interface, countries []*api.Country) []*api.Country {
	countryIdsWhereAdmin := permHelper.GetCountryIDsWithPermission("admin")
	filteredCountryMap := map[string]*api.Country{}
	for _, country := range countries {
		if countryIdsWhereAdmin.Contains(country.ID) {
			filteredCountryMap[country.ID] = country
		}
	}
	filteredCountries := make([]*api.Country, 0, len(filteredCountryMap))
	for _, country := range filteredCountryMap {
		filteredCountries = append(filteredCountries, country)
	}
	return filteredCountries
}

func parseUserForm(r *http.Request, permHelper auth.Interface, permissions *api.UserPermissions) error {
	for key, values := range r.Form {
		if strings.HasPrefix(key, "permissions[") {
			countryID := strings.TrimPrefix(key, "permissions[")
			countryID = strings.TrimSuffix(countryID, "]")
			if !permHelper.IsGlobalAdmin() && !permHelper.CanAdminCountryID(countryID) {
				return errors.New("you don't have sufficient permissions")
			}
			var permission = api.ExplicitCountryPermission{
				CountryID: countryID,
			}
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
			permissions.ExplicitCountryPermissions[countryID] = permission
		}
	}

	if len(r.Form["IsGlobalAdmin"]) != 0 {
		if !permHelper.IsGlobalAdmin() {
			return errors.New("you don't have sufficient permissions")
		}
		permissions.IsGlobalAdmin = r.FormValue("IsGlobalAdmin") == "true"
	}

	return nil
}
