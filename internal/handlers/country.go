package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func HandleCountry(templates map[string]*template.Template, repo db.CountryRepo) http.Handler {

	const (
		templateName             = "country.gohtml"
		newId                    = "new"
		pathParamCountryID       = "country_id"
		viewParamCountry         = "Country"
		formParamName            = "Name"
		formParamCode            = "Code"
		formParamNrcOrganisation = "NrcOrganisations"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx       = r.Context()
			l         = logging.NewLogger(ctx)
			err       error
			country   = &api.Country{}
			countryID = mux.Vars(r)[pathParamCountryID]
			isNew     = countryID == newId
		)

		render := func() {
			renderView(templates, templateName, w, r, viewParams{
				viewParamCountry: country,
			})
		}

		if r.Method == http.MethodGet {
			if !isNew {
				country, err = repo.GetByID(r.Context(), countryID)
				if err != nil {
					l.Error("failed to get country", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			render()
			return
		}

		if !isNew {
			country.ID = countryID
		}
		country.Name = strings.TrimSpace(r.FormValue(formParamName))
		country.Code = strings.TrimSpace(strings.ToLower(r.FormValue(formParamCode)))

		orgs := strings.Split(r.FormValue(formParamNrcOrganisation), ",")
		trimmedOrgs := []string{}
		for _, org := range orgs {
			trimmedOrgs = append(trimmedOrgs, strings.TrimSpace(org))
		}

		country.NrcOrganisations = strings.Join(trimmedOrgs, ",")

		country, err = repo.Put(r.Context(), country)
		if err != nil {
			l.Error("failed to put country", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if isNew {
			http.Redirect(w, r, "/countries/"+country.ID, http.StatusSeeOther)
		} else {
			render()
		}

	})
}
