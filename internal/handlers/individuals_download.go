package handlers

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

const queryParamFormat = "format"

func HandleDownload(
	userRepo db.IndividualRepo,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		format := r.URL.Query().Get(queryParamFormat)
		if format == "" {
			format = "xlsx"
		}

		ret, err := userRepo.GetAll(ctx, api.GetAllOptions{})
		if err != nil {
			l.Error("failed to get individuals", zap.Error(err))
			http.Error(w, "failed to get records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if format == "xlsx" {
			w.Header().Set("Content-Disposition", "attachment; filename=records.xlsx")
			w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

			if err := api.MarshalIndividualsExcel(w, ret); err != nil {
				l.Error("failed to write xlsx", zap.Error(err))
				http.Error(w, "failed to write xlsx: "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else if format == "csv" {
			w.Header().Set("Content-Disposition", "attachment; filename=records.csv")
			w.Header().Set("Content-Type", "text/csv")

			if err := api.MarshalIndividualsCSV(w, ret); err != nil {
				l.Error("failed to write csv", zap.Error(err))
				http.Error(w, "failed to write csv: "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "invalid format", http.StatusBadRequest)
			return
		}
	})
}
