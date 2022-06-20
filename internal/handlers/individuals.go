package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
)

func ListHandler(templates map[string]*template.Template, repo db.IndividualRepo, countryRepo db.CountryRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const templateName = "individuals.gohtml"
		var individuals []*api.Individual
		var getAllOptions api.GetAllOptions
		var countries []*api.Country

		render := func() {
			if err := templates[templateName].ExecuteTemplate(w, "base", map[string]interface{}{
				"Individuals": individuals,
				"Countries":   countries,
				"Options":     getAllOptions,
			}); err != nil {
				println(err.Error())
			}
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := parseGetAllOptions(r, &getAllOptions); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		errGroup, gCtx := errgroup.WithContext(r.Context())
		errGroup.Go(func() error {
			var err error
			countries, err = countryRepo.GetAll(gCtx)
			return err
		})
		errGroup.Go(func() error {
			var err error
			individuals, err = repo.GetAll(gCtx, getAllOptions)
			return err
		})
		if err := errGroup.Wait(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render()

	})
}

func parseGetAllOptions(r *http.Request, out *api.GetAllOptions) error {
	var err error
	out.Take, err = parseQryParamInt(r, "take")
	if err != nil {
		return err
	}
	if out.Take <= 0 || out.Take > 100 {
		out.Take = 20
	}

	out.Skip, err = parseQryParamInt(r, "skip")
	if err != nil {
		return err
	}
	if out.Skip < 0 {
		out.Skip = 0
	}

	out.Email = r.FormValue("email")
	out.FullName = r.FormValue("name")
	out.PhoneNumber = r.FormValue("phone_number")
	out.Address = r.FormValue("address")
	out.Genders = r.Form["gender"]
	if r.FormValue("is_minor") == "true" {
		isMinor := true
		out.IsMinor = &isMinor
	} else if r.FormValue("is_minor") == "false" {
		isMinor := false
		out.IsMinor = &isMinor
	}
	if r.FormValue("presents_protection_concerns") == "true" {
		presentsProtectionConcerns := true
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	} else if r.FormValue("presents_protection_concerns") == "false" {
		presentsProtectionConcerns := false
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	}

	ageFromStr := r.FormValue("age_from")
	if len(ageFromStr) != 0 {
		ageFrom, err := parseQryParamInt(r, "age_from")
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageFrom+1)*365)
		out.BirthDateTo = &yearsAgo
	}
	ageToStr := r.FormValue("age_to")
	if len(ageToStr) != 0 {
		ageTo, err := parseQryParamInt(r, "age_to")
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageTo+1)*365)
		out.BirthDateFrom = &yearsAgo
	}

	countries := r.FormValue("countries")
	if len(countries) != 0 {
		out.Countries = strings.Split(countries, ",")
	} else {
		out.Countries = make([]string, 0)
	}

	return nil
}