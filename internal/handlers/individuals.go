package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func ListHandler(templates map[string]*template.Template, repo db.IndividualRepo) http.Handler {

	const (
		templateName         = "individuals.gohtml"
		viewParamIndividuals = "Individuals"
		viewParamOptions     = "Options"
		queryParamCountryID  = "country_id"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			individuals   []*api.Individual
			getAllOptions api.GetAllOptions
			ctx           = r.Context()
			err           error
			l             = logging.NewLogger(ctx)
			allCountries  []*api.Country
		)

		render := func() {
			renderView(templates, templateName, w, r, viewParams{
				viewParamIndividuals: individuals,
				viewParamOptions:     getAllOptions,
			})
			return
		}

		if allCountries, err = utils.GetCountries(ctx); err != nil {
			l.Error("failed to get countries", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		countryIdMap := map[string]bool{}
		for _, c := range allCountries {
			countryIdMap[c.ID] = true
		}

		if err := r.ParseForm(); err != nil {
			l.Error("failed to parse form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := parseGetAllOptions(r, &getAllOptions); err != nil {
			l.Error("failed to parse options", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		queryCountryID := r.FormValue(queryParamCountryID)
		if queryCountryID == "" {
			// If the country ID from query param is empty, redirect to
			// the selected country ID

			if !countryIdMap[selectedCountryID] {
				// If the selected country ID is not in the list of countries,
				// return an error
				http.Error(w, "country id not found", http.StatusNotFound)
				return
			}

			url := r.URL
			q := url.Query()
			q.Set(queryParamCountryID, selectedCountryID)
			url.RawQuery = q.Encode()
			http.Redirect(w, r, url.String(), http.StatusSeeOther)
			return
		}

		if queryCountryID != selectedCountryID {
			// If the selected country ID does not match the country ID
			// from the query param, force the selection of the country ID
			// from the query param

			if !countryIdMap[queryCountryID] {
				// If the country ID from query param is not in the list of countries,
				// return an error
				http.Error(w, "country id not found", http.StatusNotFound)
				return
			}

			ctx = utils.WithSelectedCountryID(ctx, queryCountryID)
			r = r.WithContext(ctx)
			selectedCountryID = queryCountryID
		}

		authIntf, err := utils.GetAuthContext(ctx)
		if !authIntf.CanReadWriteToCountryID(selectedCountryID) {
			l.Warn("user does not have permission to read/write to country", zap.String("country_id", selectedCountryID))
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		getAllOptions.CountryID = selectedCountryID
		individuals, err = repo.GetAll(ctx, getAllOptions)
		if err != nil {
			l.Error("failed to get individuals", zap.Error(err))
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

	out.Email = r.FormValue(constants.FormParamGetIndividualsEmail)
	out.FullName = r.FormValue(constants.FormParamGetIndividualsName)
	out.PhoneNumber = r.FormValue(constants.FormParamGetIndividualsPhoneNumber)
	out.Address = r.FormValue(constants.FormParamsGetIndividualsAddress)
	out.Genders = r.Form[constants.FormParamsGetIndividualsGender]
	if r.FormValue(constants.FormParamsGetIndividualsIsMinor) == "true" {
		isMinor := true
		out.IsMinor = &isMinor
	} else if r.FormValue(constants.FormParamsGetIndividualsIsMinor) == "false" {
		isMinor := false
		out.IsMinor = &isMinor
	}
	if r.FormValue(constants.FormParamsGetIndividualsProtectionConcerns) == "true" {
		presentsProtectionConcerns := true
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	} else if r.FormValue(constants.FormParamsGetIndividualsProtectionConcerns) == "false" {
		presentsProtectionConcerns := false
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	}
	ageFromStr := r.FormValue(constants.FormParamsGetIndividualsAgeFrom)
	if len(ageFromStr) != 0 {
		ageFrom, err := parseQryParamInt(r, constants.FormParamsGetIndividualsAgeFrom)
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageFrom+1)*365)
		out.BirthDateTo = &yearsAgo
	}
	ageToStr := r.FormValue(constants.FormParamsGetIndividualsAgeTo)
	if len(ageToStr) != 0 {
		ageTo, err := parseQryParamInt(r, constants.FormParamsGetIndividualsAgeTo)
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageTo+1)*365)
		out.BirthDateFrom = &yearsAgo
	}
	out.CountryID = r.FormValue(constants.FormParamGetIndividualsCountryID)
	displacementStatuses := r.Form[constants.FormParamsGetIndividualsDisplacementStatus]
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
