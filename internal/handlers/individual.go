package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/pkg/api/deduplication"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	apivalidation "github.com/nrc-no/notcore/internal/api/validation"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/internal/views"
	apierrs "github.com/nrc-no/notcore/pkg/api/errors"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/nrc-no/notcore/pkg/views/alert"
	"github.com/nrc-no/notcore/pkg/views/bootstrap"
	"go.uber.org/zap"
)

func HandleIndividual(renderer Renderer, repo db.IndividualRepo) http.Handler {

	const (
		templateName                        = "individual.gohtml"
		templateParamAlerts                 = "Alerts"
		pathParamIndividualID               = "individual_id"
		formDeduplicationParam              = "deduplicationType"
		formParamDeduplicationLogicOperator = "deduplicationLogicOperator"
		newID                               = "new"
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
			alerts           []alert.Alert
		)

		render := func() {
			individualForm.SetErrors(validationErrors)
			renderer.RenderView(w, r, templateName, viewParams{
				"form":              individualForm,
				"Individual":        individual,
				templateParamAlerts: alerts,
			})
			return
		}

		t := locales.GetTranslator()
		successAlert := alert.Alert{
			Type:        bootstrap.StyleSuccess,
			Title:       t("success"),
			Content:     template.HTML(t("participant_saved_successfully")),
			Dismissible: true,
		}
		if r.URL.Query().Get("success") == "true" {
			alerts = append(alerts, successAlert)
		}

		// Get the currently selected Country ID
		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if !isNew {
			if individual, err = repo.GetByID(ctx, individualId); err != nil {
				l.Error("failed to get individual", zap.Error(err))
				err = apierrs.ErrorFrom(err)
				render()
				return
			}

			if individual.CountryID != selectedCountryID {
				l.Warn("user trying to access individual with the wrong country id", zap.String("individual_id", individual.ID))
				http.Error(w, fmt.Sprintf("individual not found: %v", individual.ID), http.StatusNotFound)
				return
			}
		} else {
			individual.CountryID = selectedCountryID
		}

		individualForm, err = views.NewIndividualForm(individual)
		if err != nil {
			l.Error("failed to create individual form", zap.Error(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
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

		individual.Normalize()

		warningIcon := "exclamation-triangle"

		// Validate the individual
		validationErrors = apivalidation.ValidateIndividual(individual)
		if len(validationErrors) > 0 {
			alerts = append(alerts, alert.Alert{
				Type:        bootstrap.StyleDanger,
				Title:       "Failed to save individual",
				Icon:        warningIcon,
				Content:     template.HTML("There were errors with your submission. Please correct them and try again."),
				Dismissible: true,
			})
			render()
			return
		}

		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			deduplicationTypes := r.Form[formDeduplicationParam]
			deduplicationLogicOperator := deduplication.LogicOperator(r.Form[formParamDeduplicationLogicOperator][0])
			deduplicationConfig, err := deduplication.GetDeduplicationConfig(deduplicationTypes, deduplicationLogicOperator)

			if err != nil {
				alerts = append(alerts, alert.Alert{
					Type:        bootstrap.StyleDanger,
					Title:       "An Error Occurred during Deduplication",
					Icon:        warningIcon,
					Dismissible: true,
				})
				render()
			}

			if len(deduplicationConfig.Types) > 0 {
				duplicatesInFile, duplicatesInDB, err := repo.FindDuplicates(ctx, []*api.Individual{individual}, deduplicationConfig)

				if err != nil {
					alerts = append(alerts, alert.Alert{
						Type:        bootstrap.StyleDanger,
						Title:       fmt.Sprintf("Error during deduplication %e", err),
						Icon:        warningIcon,
						Dismissible: true,
					})
					render()
					return
				}

				if duplicatesInFile != nil {
					alerts = append(alerts, alert.Alert{
						Type:        bootstrap.StyleDanger,
						Title:       fmt.Sprintf("Error during deduplication"),
						Icon:        warningIcon,
						Dismissible: true,
					})
					render()
					return
				}

				if duplicatesInDB != nil { 
					for _, dType := range deduplicationConfig.Types {
						for _, field := range dType.Config.Columns {
							value, err := individual.GetFieldValue(field)
							if err != nil || value == "" {
								continue
							}
							validationErrors = append(validationErrors, validation.Duplicate(validation.NewPath(field), value))
						}
					}

					alerts = append(alerts, alert.Alert{
						Type:        bootstrap.StyleDanger,
						Title:       fmt.Sprintf("Found %d duplicate(s)", len(duplicatesInDB)),
						Icon:        warningIcon,
						Dismissible: true,
					})
					render()
					return
				}
			}

			// Save the individual
			individual, err = repo.Put(ctx, individual, constants.IndividualDBColumns)
			if err != nil {
				l.Error("failed to put individual", zap.Error(err))
				err = apierrs.ErrorFrom(err)
				render()
				return
			}

			if individualId == "new" {
				http.Redirect(w, r, fmt.Sprintf("/countries/%s/participants/%s?success=true", individual.CountryID, individual.ID), http.StatusFound)
				return
			} else {
				if r.URL.Query().Get("success") != "true" {
					alerts = append(alerts, successAlert)
				}
				render()
				return
			}
		}
	})
}
