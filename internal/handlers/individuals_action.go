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

func HandleIndividualsAction(repo db.IndividualRepo, action string) http.Handler {

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

		var options api.ListIndividualsOptions
		if err := api.NewIndividualListFromURLValues(r.Form, &options); err != nil {
			l.Error("failed to parse options", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		individuals, err := repo.GetAll(ctx, api.ListIndividualsOptions{IDs: individualIds, Inactive: options.Inactive})
		if err != nil {
			l.Error("failed to list individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		invalidIndividualIds := validateIndividualsExistInCountry(individualIds, individuals, countryID)
		if len(invalidIndividualIds) > 0 {
			l.Warn("user trying to "+action+" individuals that don't exist or are in the wrong country", zap.Strings("individual_ids", invalidIndividualIds))
			http.Error(w, fmt.Sprintf("individuals not found: %v", invalidIndividualIds), http.StatusNotFound)
			return
		}

		if err := repo.PerformActionMany(ctx, individualIds, action); err != nil {
			l.Error("failed to "+action+" individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if options.Inactive != nil && *options.Inactive {
			http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals?inactive=true", countryID), http.StatusFound)
		} else {
			http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals", countryID), http.StatusFound)
		}
	})
}
