package handlers

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleIndividuals(renderer Renderer, repo db.IndividualRepo) http.Handler {

	const (
		templateName         = "individuals.gohtml"
		viewParamIndividuals = "Individuals"
		viewParamOptions     = "Options"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			individuals   []*api.Individual
			getAllOptions api.ListIndividualsOptions
			ctx           = r.Context()
			err           error
			l             = logging.NewLogger(ctx)
			allCountries  []*api.Country
		)

		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		render := func() {
			renderer.RenderView(w, r, templateName, map[string]interface{}{
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

		if err := api.NewIndividualListFromURLValues(r.Form, &getAllOptions); err != nil {
			l.Error("failed to parse options", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if getAllOptions.Take <= 0 || getAllOptions.Take > 100 {
			getAllOptions.Take = 20
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
