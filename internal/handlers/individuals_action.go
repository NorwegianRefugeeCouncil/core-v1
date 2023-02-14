package handlers

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func HandleIndividualsAction(renderer Renderer, repo db.IndividualRepo, action string) http.Handler {

	const (
		templateName   = "error.gohtml"
		formParamField = "individual_id"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			err error
			l   = logging.NewLogger(ctx)
		)

		renderError := func(title string, fileErrors []api.FileError) {
			renderer.RenderView(w, r, templateName, map[string]interface{}{
				"Errors": fileErrors,
				"Title":  title,
			})
		}

		countryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country", zap.Error(err))
			renderError("Failed to get selected country", nil)
			return
		}

		err = r.ParseForm()
		if err != nil {
			l.Error("failed to parse form", zap.Error(err))
			renderError("Failed to parse form", nil)
			return
		}
		individualIds := containers.NewStringSet(r.Form[formParamField]...)

		var options api.ListIndividualsOptions
		if err := api.NewIndividualListFromURLValues(r.Form, &options); err != nil {
			l.Error("failed to parse options", zap.Error(err))
			renderError("Failed to parse options", nil)
			return
		}
		options.CountryID = countryID

		individuals, err := repo.GetAll(ctx, options)
		if err != nil {
			l.Error("failed to list individuals", zap.Error(err))
			renderError("Failed to list participants", []api.FileError{{Message: "Action failed for participants", Err: []error{err}}})
			return
		}

		invalidIndividualIds := validateIndividualsExistInCountry(individualIds, individuals, countryID)
		if len(invalidIndividualIds) > 0 {
			var errors []error
			for _, individualId := range invalidIndividualIds {
				errors = append(errors, fmt.Errorf(individualId))
			}
			l.Warn("user trying to "+action+" individuals that don't exist or are in the wrong country", zap.Strings("individual_ids", invalidIndividualIds))
			renderError(fmt.Sprintf("Could not execute action %s. Please try again.", action),
				[]api.FileError{{Message: "Action failed for participants", Err: errors}})
			return
		}

		if err := repo.PerformActionMany(ctx, individualIds, action); err != nil {
			l.Error("failed to "+action+" individuals", zap.Error(err))
			renderError(fmt.Sprintf("Failed to %s participants", action), nil)
			return
		}

		r.URL.Path = fmt.Sprintf("/countries/%s/participants", countryID)
		r.Form.Del("individual_id")
		http.Redirect(w, r, r.URL.String(), http.StatusFound)

	})
}
