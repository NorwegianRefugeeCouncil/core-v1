package middleware

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func PrefetchCountries(countryRepo db.CountryRepo) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			l.Debug("country middleware")

			allCountries, err := countryRepo.GetAll(ctx)
			if err != nil {
				l.Error("failed to get all countries", zap.Error(err))
				http.Error(w, "couldn't get countries: "+err.Error(), http.StatusInternalServerError)
				return
			}

			r = r.WithContext(utils.WithCountries(ctx, allCountries))
			h.ServeHTTP(w, r)
		})
	}
}
