package handlers

import (
	"fmt"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleIndividualsDeactivate(repo db.IndividualRepo) http.Handler {

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
		individualIds := containers.NewStringSet(r.Form[formParamField]...)

		individuals, err := repo.GetAll(ctx, api.ListIndividualsOptions{IDs: individualIds})
		if err != nil {
			l.Error("failed to list individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		invalidIndividualIds := validateIndividualsExistInCountry(individualIds, individuals, countryID)
		if len(invalidIndividualIds) > 0 {
			l.Warn("user trying to deactivate individuals that don't exist or are in the wrong country", zap.Strings("individual_ids", invalidIndividualIds))
			http.Error(w, fmt.Sprintf("individuals not found: %v", invalidIndividualIds), http.StatusNotFound)
			return
		}

		if err := repo.DeactivateMany(ctx, individualIds); err != nil {
			l.Error("failed to  deactivate individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals", countryID), http.StatusFound)
	})
}