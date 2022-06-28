package handlers

import (
	"html/template"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func HandleIndividual(templates map[string]*template.Template, repo db.IndividualRepo, countryRepo db.CountryRepo) http.Handler {

	const (
		templateName              = "individual.gohtml"
		pathParamIndividualID     = "individual_id"
		newID                     = "new"
		viewParamIndividual       = "Individual"
		viewParamCountries        = "Countries"
		viewParamErrors           = "Errors"
		viewParamWasValidated     = "WasValidated"
		viewParamPermissionHelper = "PermissionHelper"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			err              error
			ctx              = r.Context()
			l                = logging.NewLogger(ctx)
			individual       = &api.Individual{}
			countries        []*api.Country
			validationErrors ValidationErrors
			wasValidated     bool
			individualId     = mux.Vars(r)[pathParamIndividualID]
			isNew            = individualId == newID
			permsHelper      permissionHelper
		)

		render := func() {
			if individual == nil {
				individual = &api.Individual{}
			}
			renderView(templates, templateName, w, r, viewParams{
				viewParamIndividual:       individual,
				viewParamCountries:        countries,
				viewParamErrors:           validationErrors,
				viewParamWasValidated:     wasValidated,
				viewParamPermissionHelper: permsHelper,
			})
			return
		}

		errGroup, gCtx := errgroup.WithContext(ctx)
		errGroup.Go(func() error {
			// Load individual
			if isNew {
				return nil
			}
			var err error
			if individual, err = repo.GetByID(gCtx, individualId); err != nil {
				l.Error("failed to get individual", zap.Error(err))
				return err
			}
			return nil
		})

		errGroup.Go(func() error {
			// Load countries
			var err error
			if countries, err = countryRepo.GetAll(gCtx); err != nil {
				l.Error("failed to get countries", zap.Error(err))
				return err
			}
			return nil
		})

		// Wait for the individual and countries to be loaded
		if err := errGroup.Wait(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		permsHelper = newPermissionHelper(ctx, countries)

		// Check if the user is allowed to read the individual
		if !isNew && !permsHelper.CanReadIndividual(individual) {
			l.Warn("user is not allowed to read individual", zap.String("individual_id", individualId))
			http.Error(w, "You are not allowed to read this individual", http.StatusForbidden)
			return
		}

		// If the user can write to a single country,
		// then pre-fill the individual with that country.
		if isNew {
			writableCountryCodes := permsHelper.GetCountryCodesWithPermission("write")
			if len(writableCountryCodes) == 1 {
				individual.Countries = append(individual.Countries, writableCountryCodes.Items()[0])
			} else {
				individual.Countries = []string{}
			}
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
		if err := parseIndividualForm(r, permsHelper, individual); err != nil {
			l.Error("failed to parse individual form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the user has permission to write to the individual
		// to the selected countries
		canWrite := permsHelper.IsGlobalAdmin()
		if !canWrite {
			for _, country := range individual.Countries {
				if permsHelper.CanWriteToCountryCode(country) {
					canWrite = true
					break
				}
			}
		}

		if !canWrite {
			l.Warn("country not allowed")
			http.Error(w, "You are not allowed to add to this country", http.StatusForbidden)
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
	if len(individual.Countries) == 0 {
		ret[formParamIndividualCountries] = "At least one country is required"
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
	for i, c := range individual.Countries {
		individual.Countries[i] = trimString(c)
	}
	sort.Strings(individual.Countries)
}

func parseIndividualForm(r *http.Request, permsHelper permissionHelper, individual *api.Individual) error {
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

	wantCountryCodes := containers.NewStringSet()
	for _, c := range strings.Split(r.FormValue(formParamIndividualCountries), ",") {
		c = trimString(c)
		if c != "" {
			wantCountryCodes.Add(c)
		}
	}

	if permsHelper.IsGlobalAdmin() {
		individual.Countries = wantCountryCodes.Items()
	} else {
		previousCountryCodes := containers.NewStringSet(individual.Countries...)
		writableCountryCodes := permsHelper.GetCountryCodesWithPermission("write")

		finalCountryCodes := containers.NewStringSet()
		finalCountryCodes.Add(previousCountryCodes.Difference(writableCountryCodes).Items()...)
		finalCountryCodes.Add(wantCountryCodes.Intersection(writableCountryCodes).Items()...)
		individual.Countries = finalCountryCodes.Items()
	}

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
		case csvHeaderIndividualCountries:
			individual.Countries = strings.Split(cols[idx], ",")
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
