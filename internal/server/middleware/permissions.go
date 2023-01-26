package middleware

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

// ParsedPermissions is a helper struct to store the parsed permissions
type ParsedPermissions struct {
	IsGlobalAdmin bool
	CountryIds    containers.StringSet
	CanRead       bool
	CanWrite      bool
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

			perms := parsePermissions(allCountries, jwtGroups, session.GetUserGroups(), session.GetNrcOrganisation())
			authIntf := auth.New(perms.CountryIds, allCountryIDs, perms.IsGlobalAdmin, perms.CanRead, perms.CanWrite)
			r = r.WithContext(utils.WithAuthContext(ctx, authIntf))
			h.ServeHTTP(w, r)

		})
	}
}

// parsePermissions will retrieve the country ids from the user's groups
// and determine if the user is a global admin
func parsePermissions(allCountries []*api.Country, jwtGroups utils.JwtGroupOptions, userGroups []string, nrcOrganisation string) *ParsedPermissions {
	countryIds := containers.NewStringSet()

	for _, c := range allCountries {
		for _, org := range c.NrcOrganisations.Items() {
			if nrcOrganisation == org {
				countryIds.Add(c.ID)
			}
		}
	}

	isGlobalAdmin := false
	canRead := false
	canWrite := false
	for _, group := range userGroups {
		if group == jwtGroups.GlobalAdmin {
			isGlobalAdmin = true
			canWrite = true
			canRead = true
			continue
		}
		if group == jwtGroups.CanWrite {
			canWrite = true
			canRead = true
			continue
		}
		if group == jwtGroups.CanRead {
			canRead = true
			continue
		}
	}

	return &ParsedPermissions{
		IsGlobalAdmin: isGlobalAdmin,
		CanWrite:      canWrite,
		CanRead:       canRead,
		CountryIds:    countryIds,
	}
}
