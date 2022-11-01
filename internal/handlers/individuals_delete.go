package handlers

import (
	"fmt"
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleIndividualsDelete(repo db.IndividualRepo) http.Handler {

	const (
		formParamField = "individual_id"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			err error
			l   = logging.NewLogger(ctx)
		)

		countryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = r.ParseForm()
		if err != nil {
			l.Error("failed to parse form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		individualIds := r.Form[formParamField]

		individuals, err := repo.ListByID(ctx, individualIds)
		if err != nil {
			l.Error("failed to list individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, individual := range individuals {
			if individual.CountryID != countryID {
				l.Error("individual does not belong to selected country", zap.String("individual_id", individual.ID), zap.String("country_id", countryID))
				http.Error(w, fmt.Sprintf("individual %s does not belong to selected country", individual.ID), http.StatusBadRequest)
				return
			}
		}

		if err := repo.SoftDeleteMany(ctx, individualIds); err != nil {
			l.Error("failed to delete individual", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals", countryID), http.StatusFound)
	})
}
