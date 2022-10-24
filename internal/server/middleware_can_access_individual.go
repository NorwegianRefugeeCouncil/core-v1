package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func canAccessIndividualMiddleware(repo db.IndividualRepo) func(handler http.Handler) http.Handler {
	const (
		pathParamIndividualID = "individual_id"
		newID                 = "new"
	)

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				err          error
				ctx          = r.Context()
				l            = logging.NewLogger(ctx)
				individual   = &api.Individual{}
				individualId = mux.Vars(r)[pathParamIndividualID]
				isNew        = individualId == newID
			)

			if !isNew {
				if individual, err = repo.GetByID(ctx, individualId); err != nil {
					l.Error("failed to get individual", zap.Error(err))
					http.Error(w, "failed to get individual", http.StatusInternalServerError)
				}
			}

			authInterface, err := utils.GetAuthContext(ctx)
			if err != nil {
				l.Error("failed to get auth context", zap.Error(err))
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			if !isNew && !authInterface.CanReadWriteToCountryID(individual.CountryID) {
				if r.Method == http.MethodGet {
					if !authInterface.CanReadWriteToCountryID(individual.CountryID) {
						l.Warn("user is not allowed to read individual", zap.String("individual_id", individualId))
						http.Error(w, "forbidden", http.StatusForbidden)
						return
					}
				} else if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
					if !authInterface.CanReadWriteToCountryID(individual.CountryID) {
						l.Warn("user is not allowed to write to individual", zap.String("individual_id", individualId))
						http.Error(w, "forbidden", http.StatusForbidden)
						return
					}
				} else {
					l.Warn("method not allowed", zap.String("method", r.Method))
					http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
					return
				}
			}

			h.ServeHTTP(w, r)
		})
	}
}
