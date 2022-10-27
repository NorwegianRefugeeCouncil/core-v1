package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func SelectedCountry() func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		const pathParamCountryID = "country_id"

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			countryId := mux.Vars(r)[pathParamCountryID]

			allCountries, err := utils.GetCountries(ctx)
			if err != nil {
				l.Error("failed to get all countries", zap.Error(err))
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			// checking if the selected country id is actually a valid country id
			found := false
			for _, c := range allCountries {
				if c.ID == countryId {
					found = true
					break
				}
			}

			if found {
				// Store the selected country ID in the request context
				r = r.WithContext(utils.WithSelectedCountryID(ctx, countryId))
			}
			h.ServeHTTP(w, r)
		})
	}
}
