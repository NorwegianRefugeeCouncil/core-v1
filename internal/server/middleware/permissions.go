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

// permissionMiddleware will compute the user's permissions and add them to the context
func ComputePermissions(
	globalAdminGroup string,
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

			perms := parsePermissions(allCountries, globalAdminGroup, session.GetUserGroups(), session.GetNrcOrganisation())
			authIntf := auth.New(perms.countryIds, allCountryIDs, perms.isGlobalAdmin)
			r = r.WithContext(utils.WithAuthContext(ctx, authIntf))
			h.ServeHTTP(w, r)

		})
	}
}

// parsedPermissions is a helper struct to store the parsed permissions
type parsedPermissions struct {
	isGlobalAdmin bool
	countryIds    containers.StringSet
}

// parsePermissions will retrieve the country ids from the user's groups
// and determine if the user is a global admin
func parsePermissions(allCountries []*api.Country, globalAdminGroup string, userGroups []string, nrcOrganisation string) *parsedPermissions {
	countryIds := containers.NewStringSet()

	// maps a jwt group name to a country id
	countryGroupMap := make(map[string]string)
	for _, c := range allCountries {
		if c.NrcOrganisation == "" {
			continue
		}
		countryGroupMap[c.NrcOrganisation] = c.ID
	}

	isGlobalAdmin := false
	for _, group := range userGroups {
		if group == globalAdminGroup {
			isGlobalAdmin = true
			continue
		}
	}

	if countryID, ok := countryGroupMap[nrcOrganisation]; ok {
		countryIds.Add(countryID)
	}

	return &parsedPermissions{
		isGlobalAdmin: isGlobalAdmin,
		countryIds:    countryIds,
	}
}
