package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/nrc-no/notcore/internal/api"
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
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			individuals   []*api.Individual
			getAllOptions api.GetAllOptions
			ctx           = r.Context()
			err           error
			l             = logging.NewLogger(ctx)
		)

		render := func() {
			renderView(templates, templateName, w, r, viewParams{
				viewParamIndividuals: individuals,
				viewParamOptions:     getAllOptions,
			})
			return
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
	out.CountryID = r.FormValue(formParamGetIndividualsCountryID)
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
