package middleware

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

// ParsedPermissions is a helper struct to store the parsed permissions
type ParsedPermissions struct {
	IsGlobalAdmin bool
	CountryPermissions auth.CountryPermissions 
}

// permissionMiddleware will compute the user's permissions and add them to the context
func ComputePermissions(
	jwtGroups utils.JwtGroupOptions,
) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			session, ok := utils.GetSession(ctx)
			if !ok {
				l.Error("no session found in context")
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			allCountries, err := utils.GetCountries(ctx)
			if err != nil {
				l.Error("failed to get all countries", zap.Error(err))
				http.Error(w, "couldn't get countries: "+err.Error(), http.StatusInternalServerError)
				return
			}

			allCountryIDs := containers.NewStringSet()
			for _, c := range allCountries {
				allCountryIDs.Add(c.ID)
			}

			perms := parsePermissions(allCountries, jwtGroups, session.GetUserGroups())
			authIntf := auth.New(perms.CountryPermissions, allCountryIDs, perms.IsGlobalAdmin)
			r = r.WithContext(utils.WithAuthContext(ctx, authIntf))
			h.ServeHTTP(w, r)

		})
	}
}

// parsePermissions will retrieve the country ids from the user's groups
// and determine if the user is a global admin
func parsePermissions(allCountries []*api.Country, jwtGroups utils.JwtGroupOptions, userGroups []string) *ParsedPermissions {
	countryPermissions := auth.CountryPermissions{}
	userGroupsSet := containers.NewStringSet(userGroups...)

	for _, c := range allCountries {
		if userGroupsSet.Contains(c.ReadGroup) {
			if countryPermissions[c.ID] == nil {
				countryPermissions[c.ID] = containers.NewSet[auth.Permission]()
			}
			countryPermissions[c.ID].Add(auth.PermissionRead)
		}

		if userGroupsSet.Contains(c.WriteGroup) {
			if countryPermissions[c.ID] == nil {
				countryPermissions[c.ID] = containers.NewSet[auth.Permission]()
			}
			countryPermissions[c.ID].Add(auth.PermissionWrite)
		}
	}

	isGlobalAdmin := userGroupsSet.Contains(jwtGroups.GlobalAdmin)

	return &ParsedPermissions{
		IsGlobalAdmin: isGlobalAdmin,
		CountryPermissions: countryPermissions,
	}
}
