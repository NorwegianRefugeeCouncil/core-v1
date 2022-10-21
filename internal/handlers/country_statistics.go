package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleCountryStatistics(repo db.StatisticRepo) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
			err error
		)

		authIntf, err := utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		countryID := mux.Vars(r)["country_id"]
		if !authIntf.CanReadWriteToCountryID(countryID) {
			l.Warn("cannot access country statistics")
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		stats, err := repo.CountryStatistics(ctx, countryID)
		if err != nil {
			l.Error("failed to get country statistics", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(stats)

	})
}
