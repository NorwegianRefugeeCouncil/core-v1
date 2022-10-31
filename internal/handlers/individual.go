package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	apivalidation "github.com/nrc-no/notcore/internal/api/validation"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/internal/views"
	apierrs "github.com/nrc-no/notcore/pkg/api/errors"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/nrc-no/notcore/pkg/views/alerts"
	"go.uber.org/zap"
)

func HandleIndividual(templates map[string]*template.Template, repo db.IndividualRepo) http.Handler {

	const (
		templateName          = "individual.gohtml"
		pathParamIndividualID = "individual_id"
		newID                 = "new"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			err              error
			ctx              = r.Context()
			l                = logging.NewLogger(ctx)
			individual       = &api.Individual{}
			validationErrors validation.ErrorList
			individualId     = mux.Vars(r)[pathParamIndividualID]
			isNew            = individualId == newID
			individualForm   *views.IndividualForm
			viewAlerts       = alerts.NewAlerts()
		)

		render := func() {
			individualForm.SetErrors(validationErrors)
			renderView(templates, templateName, w, r, viewParams{
				"Form":       individualForm,
				"Individual": individual,
				"Alerts":     viewAlerts,
			})
			return
		}

		if !isNew {
			if individual, err = repo.GetByID(ctx, individualId); err != nil {
				l.Error("failed to get individual", zap.Error(err))
				err = apierrs.ErrorFrom(err)
				render()
				return
			}
		}

		// Get the currently selected Country ID
		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			err = apierrs.ErrorFrom(err)
			render()
			return
		}

		individual.CountryID = selectedCountryID
		individualForm, err = views.NewIndividualForm(individual)
		if err != nil {
			l.Error("failed to create individual form", zap.Error(err))
			err = apierrs.ErrorFrom(err)
			render()
			return
		}

		// Render the form if GET
		if r.Method == http.MethodGet {
			render()
			return
		}

		// Parse the form
		if err := r.ParseForm(); err != nil {
			l.Error("failed to parse form", zap.Error(err))
			err = apierrs.ErrorFrom(err)
			render()
			return
		}

		individualForm.ParseURLValues(r.Form)
		if err := individualForm.Into(individual); err != nil {
			l.Error("failed to parse individual form", zap.Error(err))
			err = apierrs.ErrorFrom(err)
			render()
			return
		}
		individual.CountryID = selectedCountryID
		if !isNew {
			individual.ID = individualId
		} else {
			individual.ID = uuid.New().String()
		}

		// Validate the individual
		validationErrors = apivalidation.ValidateIndividual(individual)
		if len(validationErrors) > 0 {
			viewAlerts = append(viewAlerts, &alerts.Alert{
				Style:       alerts.AlertStyleDanger,
				Title:       "Validation errors",
				Body:        "Please correct the errors below",
				Dismissible: true,
				Icon:        "exclamation-triangle",
			})
			render()
			return
		}

		// Save the individual
		individual, err = repo.Put(ctx, individual, constants.IndividualDBColumns)
		if err != nil {
			l.Error("failed to put individual", zap.Error(err))
			err = apierrs.ErrorFrom(err)
			render()
			return
		}

		viewAlerts = append(viewAlerts, &alerts.Alert{
			Style:       alerts.AlertStyleSuccess,
			Body:        fmt.Sprintf("Individual %s saved", individual.FullName),
			Dismissible: true,
			Icon:        "check-circle",
		})

		if isNew {
			http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals/%s", individual.CountryID, individual.ID), http.StatusFound)
			return
		} else {
			render()
			return
		}

	})
}
