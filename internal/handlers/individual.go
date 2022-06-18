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
	"github.com/nrc-no/notcore/internal/utils"
)

func HandleIndividual(templates map[string]*template.Template, repo db.IndividualRepo, countryRepo db.CountryRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var individual = &api.Individual{}

		individualId := mux.Vars(r)["individual_id"]

		ctx := r.Context()

		countries, err := countryRepo.GetAll(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render := func() {
			if individual == nil {
				individual = &api.Individual{}
			}
			if err := templates["individual.gohtml"].ExecuteTemplate(w, "base", map[string]interface{}{
				"Individual": individual,
				"Countries":  countries,
			}); err != nil {
				println(err.Error())
			}
			return
		}

		individualId = mux.Vars(r)["individual_id"]
		if individualId != "new" {
			individual, err = repo.GetByID(ctx, individualId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if r.Method == "GET" {
			render()
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := parseIndividualForm(r, individual); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = repo.Put(ctx, individual, api.AllndividualFields)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/individuals/"+individual.ID, http.StatusSeeOther)

	})
}

func normalizeIndividual(individual *api.Individual) {
	individual.FullName = trimString(individual.FullName)
	individual.PreferredName = trimString(individual.PreferredName)
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
	individual.FullName = r.FormValue("FullName")
	individual.PreferredName = r.FormValue("PreferredName")
	individual.DisplacementStatus = r.FormValue("DisplacementStatus")
	individual.Email = r.FormValue("Email")
	individual.PhoneNumber = r.FormValue("PhoneNumber")
	individual.Address = r.FormValue("Address")
	individual.Gender = r.FormValue("Gender")
	individual.BirthDate, err = parseBirthDate(r.FormValue("BirthDate"))
	if err != nil {
		return err
	}
	individual.IsMinor = r.FormValue("IsMinor") == "true"
	individual.PresentsProtectionConcerns = r.FormValue("PresentsProtectionConcerns") == "true"
	individual.PhysicalImpairment = r.FormValue("PhysicalImpairment")
	individual.MentalImpairment = r.FormValue("MentalImpairment")
	individual.SensoryImpairment = r.FormValue("SensoryImpairment")
	individual.Countries = strings.Split(r.FormValue("Countries"), ",")
	normalizeIndividual(individual)
	return nil
}

func parseIndividualCsvRow(colMapping map[string]int, cols []string) (*api.Individual, error) {
	var err error
	var individual = &api.Individual{}
	for field, idx := range colMapping {
		switch field {
		case "id":
			individual.ID = cols[idx]
		case "full_name":
			individual.FullName = cols[idx]
		case "preferred_name":
			individual.PreferredName = cols[idx]
		case "displacement_status":
			individual.DisplacementStatus = cols[idx]
		case "phone_number":
			individual.PhoneNumber = cols[idx]
		case "email":
			individual.Email = cols[idx]
		case "address":
			individual.Address = cols[idx]
		case "gender":
			individual.Gender = cols[idx]
		case "birth_date":
			individual.BirthDate, err = parseBirthDate(cols[idx])
			if err != nil {
				return nil, err
			}
		case "is_minor":
			individual.IsMinor = cols[idx] == "true"
		case "presents_protection_concerns":
			individual.PresentsProtectionConcerns = cols[idx] == "true"
		case "physical_impairment":
			individual.PhysicalImpairment = cols[idx]
		case "sensory_impairment":
			individual.SensoryImpairment = cols[idx]
		case "mental_impairment":
			individual.MentalImpairment = cols[idx]
		case "countries":
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
