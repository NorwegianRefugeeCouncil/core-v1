package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/internal/validation"
	"go.uber.org/zap"
)

func HandleIndividual(templates map[string]*template.Template, repo db.IndividualRepo) http.Handler {

	const (
		templateName          = "individual.gohtml"
		pathParamIndividualID = "individual_id"
		newID                 = "new"
		viewParamIndividual   = "Individual"
		viewParamErrors       = "Errors"
		viewParamWasValidated = "WasValidated"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			err              error
			ctx              = r.Context()
			l                = logging.NewLogger(ctx)
			individual       = &api.Individual{}
			validationErrors validation.ValidationErrors
			wasValidated     bool
			individualId     = mux.Vars(r)[pathParamIndividualID]
			isNew            = individualId == newID
			authIntf         auth.Interface
		)

		render := func() {
			if individual == nil {
				individual = &api.Individual{}
			}
			renderView(templates, templateName, w, r, viewParams{
				viewParamIndividual:   individual,
				viewParamErrors:       validationErrors,
				viewParamWasValidated: wasValidated,
			})
			return
		}

		if !isNew {
			if individual, err = repo.GetByID(ctx, individualId); err != nil {
				l.Error("failed to get individual", zap.Error(err))
				http.Error(w, "failed to get individual", http.StatusInternalServerError)
			}
		}

		authIntf, err = utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Check if the user is allowed to read the individual
		if !isNew && !authIntf.CanReadWriteToCountryID(individual.CountryID) {
			l.Warn("user is not allowed to read individual", zap.String("individual_id", individualId))
			http.Error(w, "You are not allowed to read this individual", http.StatusForbidden)
			return
		}

		// Get the currently selected Country ID
		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse the individual form
		if err := parseIndividualForm(r, authIntf, individual); err != nil {
			l.Error("failed to parse individual form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		individual.CountryID = selectedCountryID

		// Check if the user has permission to write to the country
		if !authIntf.CanReadWriteToCountryID(individual.CountryID) {
			l.Warn("user is not allowed to create an individual for country",
				zap.String("country_id", individual.CountryID))
			http.Error(w, "You are not allowed to add an individual to this country", http.StatusForbidden)
			return
		}

		// Validate the individual
		validationErrors = validation.ValidateIndividual(individual)
		if len(validationErrors) > 0 {
			render()
			return
		}

		// Save the individual
		individual, err = repo.Put(ctx, individual, constants.IndividualDBColumns)
		if err != nil {
			l.Error("failed to put individual", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		wasValidated = true
		if individualId == "new" {
			http.Redirect(w, r, "/individuals/"+individual.ID, http.StatusFound)
			return
		} else {
			render()
			return
		}

	})
}

func parseIndividualForm(r *http.Request, permsHelper auth.Interface, individual *api.Individual) error {
	var err error
	individual.FullName = r.FormValue(constants.FormParamIndividualFullName)
	individual.PreferredName = r.FormValue(constants.FormParamIndividualPreferredName)
	individual.DisplacementStatus = r.FormValue(constants.FormParamIndividualDisplacementStatus)
	individual.Email = r.FormValue(constants.FormParamIndividualEmail)
	individual.PhoneNumber = r.FormValue(constants.FormParamIndividualPhoneNumber)
	individual.Address = r.FormValue(constants.FormParamIndividualAddress)
	individual.Gender = r.FormValue(constants.FormParamIndividualGender)
	individual.BirthDate, err = api.ParseDate(r.FormValue(constants.FormParamIndividualBirthDate))
	if err != nil {
		return err
	}
	individual.IsMinor = r.FormValue(constants.FormParamIndividualIsMinor) == "true"
	individual.PresentsProtectionConcerns = r.FormValue(constants.FormParamIndividualPresentsProtectionConcerns) == "true"
	individual.PhysicalImpairment = r.FormValue(constants.FormParamIndividualPhysicalImpairment)
	individual.MentalImpairment = r.FormValue(constants.FormParamIndividualMentalImpairment)
	individual.SensoryImpairment = r.FormValue(constants.FormParamIndividualSensoryImpairment)
	individual.CountryID = r.FormValue(constants.FormParamIndividualCountry)
	individual.Normalize()
	return nil
}
