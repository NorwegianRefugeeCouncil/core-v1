package handlers

import (
	"fmt"
	"net/http"

<<<<<<< HEAD
	"github.com/nrc-no/notcore/internal/api"
=======
>>>>>>> 3cca0ba (Add handler to soft delete many individuals)
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

		individualIds := r.MultipartForm.Value[formParamField]

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

		individuals, err := repo.GetAll(ctx, api.GetAllOptions{IDs: individualIds})
		if err != nil {
			l.Error("failed to list individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var invalidIndividualIds []string
		for _, individual := range individuals {
			if individual.CountryID != countryID {
				invalidIndividualIds = append(invalidIndividualIds, individual.ID)
			}
		}
		if len(invalidIndividualIds) > 0 {
			l.Error("individuals do not belong to selected country", zap.Strings("individual_ids", invalidIndividualIds))
			http.Error(w, fmt.Sprintf("individuals not found: %v", invalidIndividualIds), http.StatusNotFound)
			return
		}

		if err := repo.SoftDeleteMany(ctx, individualIds); err != nil {
			l.Error("failed to delete individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals", countryID), http.StatusFound)
	})
}
