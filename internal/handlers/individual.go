package handlers

import (
	"html/template"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func HandleIndividual(templates map[string]*template.Template, repo db.IndividualRepo, countryRepo db.CountryRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const (
			templateName          = "individual.gohtml"
			pathParamIndividualID = "individual_id"
			newID                 = "new"
			viewParamIndividual   = "Individual"
			viewParamCountries    = "Countries"
			viewParamErrors       = "Errors"
			viewParamWasValidated = "WasValidated"
		)

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
		)

		render := func() {
			if individual == nil {
				individual = &api.Individual{}
			}
			renderView(templates, templateName, w, r, map[string]interface{}{
				viewParamIndividual:   individual,
				viewParamCountries:    countries,
				viewParamErrors:       validationErrors,
				viewParamWasValidated: wasValidated,
			})
			return
		}

		errGroup, gCtx := errgroup.WithContext(ctx)
		errGroup.Go(func() error {
			var err error
			if !isNew {
				individual, err = repo.GetByID(gCtx, individualId)
				if err != nil {
					l.Error("failed to get individual", zap.Error(err))
					return err
				}
			}
			return nil
		})
		errGroup.Go(func() error {
			var err error
			countries, err = countryRepo.GetAll(gCtx)
			if err != nil {
				l.Error("failed to get countries", zap.Error(err))
				return err
			}
			return nil
		})
		if err := errGroup.Wait(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Method == http.MethodGet {
			render()
			return
		}

		if err := r.ParseForm(); err != nil {
			l.Error("failed to parse form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := parseIndividualForm(r, individual); err != nil {
			l.Error("failed to parse individual form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		validationErrors = validateIndividual(individual)
		if len(validationErrors) > 0 {
			render()
			return
		}

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

type ValidationErrors map[string]string

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

func parseIndividualForm(r *http.Request, individual *api.Individual) error {

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
	individual.Countries = make([]string, 0)
	countries := strings.Split(r.FormValue(formParamIndividualCountries), ",")
	for _, c := range countries {
		c = trimString(c)
		if c != "" {
			individual.Countries = append(individual.Countries, c)
		}
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
