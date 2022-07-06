package handlers

import (
	"github.com/nrc-no/notcore/internal/clients/zanzibar"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func ListHandler(templates map[string]*template.Template, client *zanzibar.ZanzibarClient, repo db.IndividualRepo, countryRepo db.CountryRepo) http.Handler {

	const (
		templateName                = "individuals.gohtml"
		viewParamIndividuals        = "Individuals"
		viewParamCountries          = "Countries"
		viewParamOptions            = "Options"
		viewParamsPermissionsHelper = "PermissionsHelper"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			individuals   []*api.Individual
			getAllOptions api.GetAllOptions
			countries     []*api.Country
			ctx           = r.Context()
			l             = logging.NewLogger(ctx)
			err           error
			permsHelper   permissionHelper
		)

		render := func() {
			renderView(templates, templateName, w, r, viewParams{
				viewParamIndividuals:        individuals,
				viewParamCountries:          countries,
				viewParamOptions:            getAllOptions,
				viewParamsPermissionsHelper: permsHelper,
			})
			return
		}

		if err := r.ParseForm(); err != nil {
			logging.NewLogger(ctx).Error("failed to parse form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := parseGetAllOptions(r, &getAllOptions); err != nil {
			logging.NewLogger(ctx).Error("failed to parse options", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		countries, err = countryRepo.GetAll(ctx)
		if err != nil {
			l.Error("failed to get countries", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		permsHelper = newPermissionHelper(ctx, countries)
		getAllOptionsForQry := applyPermissionsToIndividualsQuery(permsHelper, getAllOptions)

		individuals, err = repo.GetAll(ctx, getAllOptionsForQry)
		if err != nil {
			l.Error("failed to get individuals", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render()

	})
}

func applyPermissionsToIndividualsQuery(
	permsHelper permissionHelper,
	getAllOptions api.GetAllOptions,
) api.GetAllOptions {
	if permsHelper.isGlobalAdmin {
		return getAllOptions
	}
	allowedCountryCodes := permsHelper.GetCountryCodesWithAnyPermission("read", "write", "admin")
	if getAllOptions.Countries == nil {
		getAllOptions.Countries = allowedCountryCodes.Items()
		return getAllOptions
	}
	wantCountryCodes := containers.NewStringSet(getAllOptions.Countries...)
	finalCountryCodes := wantCountryCodes.Intersection(allowedCountryCodes)
	getAllOptions.Countries = finalCountryCodes.Items()
	return getAllOptions
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

	out.Email = r.FormValue(formParamGetIndividualsEmail)
	out.FullName = r.FormValue(formParamGetIndividualsName)
	out.PhoneNumber = r.FormValue(formParamGetIndividualsPhoneNumber)
	out.Address = r.FormValue(formParamsGetIndividualsAddress)
	out.Genders = r.Form[formParamsGetIndividualsGender]
	if r.FormValue(formParamsGetIndividualsIsMinor) == "true" {
		isMinor := true
		out.IsMinor = &isMinor
	} else if r.FormValue(formParamsGetIndividualsIsMinor) == "false" {
		isMinor := false
		out.IsMinor = &isMinor
	}
	if r.FormValue(formParamsGetIndividualsProtectionConcerns) == "true" {
		presentsProtectionConcerns := true
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	} else if r.FormValue(formParamsGetIndividualsProtectionConcerns) == "false" {
		presentsProtectionConcerns := false
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	}
	ageFromStr := r.FormValue(formParamsGetIndividualsAgeFrom)
	if len(ageFromStr) != 0 {
		ageFrom, err := parseQryParamInt(r, formParamsGetIndividualsAgeFrom)
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageFrom+1)*365)
		out.BirthDateTo = &yearsAgo
	}
	ageToStr := r.FormValue(formParamsGetIndividualsAgeTo)
	if len(ageToStr) != 0 {
		ageTo, err := parseQryParamInt(r, formParamsGetIndividualsAgeTo)
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageTo+1)*365)
		out.BirthDateFrom = &yearsAgo
	}
	countries := r.FormValue(formParamGetIndividualsCountries)
	if len(countries) != 0 {
		countrySet := containers.NewStringSet()
		for _, c := range strings.Split(countries, ",") {
			c = trimString(c)
			if len(c) == 0 {
				continue
			}
			countrySet.Add(c)
		}
		out.Countries = countrySet.Items()
	} else {
		out.Countries = nil
	}
	displacementStatuses := r.Form[formParamsGetIndividualsDisplacementStatus]
	var displacementStatusMap = map[string]bool{}
	for _, s := range displacementStatuses {
		if displacementStatusMap[s] {
			continue
		}
		displacementStatusMap[s] = true
		out.DisplacementStatuses = append(out.DisplacementStatuses, s)
	}
	return nil
}
