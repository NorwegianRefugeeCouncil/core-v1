package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
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
			validationErrors ValidationErrors
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
		if !isNew && !authIntf.CanReadFromCountryID(individual.CountryID) {
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

		if !isNew {
			// At this point, the user might have followed a link to an existing individual.
			// At this point, the user might have followed a link to an existing individual.
			// That individual might not be in the users' selected country.
			// In this case, update the users' selected country to match the individual
			ctx = utils.WithSelectedCountryID(ctx, individual.CountryID)
			r = r.WithContext(ctx)
			setSelectedCountryCookie(w, individual.CountryID)
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
		if !authIntf.CanWriteToCountryID(individual.CountryID) {
			l.Warn("user is not allowed to create an individual for country",
				zap.String("country_id", individual.CountryID))
			http.Error(w, "You are not allowed to add an individual to this country", http.StatusForbidden)
			return
		}

		// Validate the individual
		validationErrors = validateIndividual(individual)
		if len(validationErrors) > 0 {
			render()
			return
		}

		// Save the individual
		individual, err = repo.Put(ctx, individual, api.AllndividualFields)
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

func validateIndividual(individual *api.Individual) ValidationErrors {
	ret := ValidationErrors{}
	if individual.FullName == "" {
		ret[formParamIndividualFullName] = "Full name is required"
	}
	if len(individual.CountryID) == 0 {
		ret[formParamIndividualCountry] = "Country is required"
	}
	return ret
}

func normalizeIndividual(individual *api.Individual) {
	individual.FullName = trimString(individual.FullName)
	individual.PreferredName = trimString(individual.PreferredName)
	if individual.PreferredName == "" {
		individual.PreferredName = individual.FullName
	}
	individual.DisplacementStatus = trimString(individual.DisplacementStatus)
	individual.Email = trimString(utils.NormalizeEmail(individual.Email))
	individual.PhoneNumber = trimString(individual.PhoneNumber)
	individual.Address = trimString(individual.Address)
	individual.Gender = trimString(individual.Gender)
	individual.NormalizedPhoneNumber = utils.NormalizePhoneNumber(individual.PhoneNumber)
	individual.PhysicalImpairment = trimString(individual.PhysicalImpairment)
	individual.MentalImpairment = trimString(individual.MentalImpairment)
	individual.SensoryImpairment = trimString(individual.SensoryImpairment)
}

func parseIndividualForm(r *http.Request, permsHelper auth.Interface, individual *api.Individual) error {
	var err error
	individual.FullName = r.FormValue(formParamIndividualFullName)
	individual.PreferredName = r.FormValue(formParamIndividualPreferredName)
	individual.DisplacementStatus = r.FormValue(formParamIndividualDisplacementStatus)
	individual.Email = r.FormValue(formParamIndividualEmail)
	individual.PhoneNumber = r.FormValue(formParamIndividualPhoneNumber)
	individual.Address = r.FormValue(formParamIndividualAddress)
	individual.Gender = r.FormValue(formParamIndividualGender)
	individual.BirthDate, err = parseBirthDate(r.FormValue(formParamIndividualBirthDate))
	if err != nil {
		return err
	}
	individual.IsMinor = r.FormValue(formParamIndividualIsMinor) == "true"
	individual.PresentsProtectionConcerns = r.FormValue(formParamIndividualPresentsProtectionConcerns) == "true"
	individual.PhysicalImpairment = r.FormValue(formParamIndividualPhysicalImpairment)
	individual.MentalImpairment = r.FormValue(formParamIndividualMentalImpairment)
	individual.SensoryImpairment = r.FormValue(formParamIndividualSensoryImpairment)
	individual.CountryID = r.FormValue(formParamIndividualCountry)
	normalizeIndividual(individual)
	return nil
}

func parseIndividualCsvRow(colMapping map[string]int, cols []string) (*api.Individual, error) {
	var err error
	var individual = &api.Individual{}
	for field, idx := range colMapping {
		switch field {
		case csvHeaderIndividualID:
			individual.ID = cols[idx]
		case csvHeaderIndividualFullName:
			individual.FullName = cols[idx]
		case csvHeaderIndividualPreferredName:
			individual.PreferredName = cols[idx]
		case csvHeaderIndividualDisplacementStatus:
			individual.DisplacementStatus = cols[idx]
		case csvHeaderIndividualPhoneNumber:
			individual.PhoneNumber = cols[idx]
		case csvHeaderIndividualEmail:
			individual.Email = cols[idx]
		case csvHeaderIndividualAddress:
			individual.Address = cols[idx]
		case csvHeaderIndividualGender:
			individual.Gender = cols[idx]
		case csvHeaderIndividualBirthDate:
			individual.BirthDate, err = parseBirthDate(cols[idx])
			if err != nil {
				return nil, err
			}
		case csvHeaderIndividualIsMinor:
			individual.IsMinor = cols[idx] == "true"
		case csvHeaderIndividualPresentsProtectionConcerns:
			individual.PresentsProtectionConcerns = cols[idx] == "true"
		case csvHeaderIndividualPhysicalImpairment:
			individual.PhysicalImpairment = cols[idx]
		case csvHeaderIndividualSensoryImpairment:
			individual.SensoryImpairment = cols[idx]
		case csvHeaderIndividualMentalImpairment:
			individual.MentalImpairment = cols[idx]
		case csvHeaderIndividualCountryID:
			individual.CountryID = cols[idx]
		}

	}
	normalizeIndividual(individual)
	return individual, nil
}

func parseBirthDate(s string) (*time.Time, error) {
	if s != "" {
		birthDate, err := time.Parse("2006-01-02", s)
		if err != nil {
			return nil, err
		}
		return &birthDate, nil
	}
	return nil, nil
}

func trimString(s string) string {
	return strings.Trim(s, " \t\n\r")
}
