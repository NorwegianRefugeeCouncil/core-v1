package handlers

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func HandleDownload(
	userRepo db.IndividualRepo,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		ret, err := userRepo.GetAll(ctx, api.GetAllOptions{})
		if err != nil {
			l.Error("failed to get individuals", zap.Error(err))
			http.Error(w, "failed to get records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=records.csv")
		w.Header().Set("Content-Type", "text/csv")

		if err := api.MarshalIndividualsCSV(w, ret); err != nil {
			l.Error("failed to write csv", zap.Error(err))
			http.Error(w, "failed to write csv: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
