package handlers

import (
	"github.com/nrc-no/notcore/internal/clients/zanzibar"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleCountry(templates map[string]*template.Template, client *zanzibar.ZanzibarClient, repo db.CountryRepo) http.Handler {

	const (
		templateName       = "country.gohtml"
		newId              = "new"
		pathParamCountryID = "country_id"
		viewParamCountry   = "Country"
		formParamName      = "Name"
		formParamCode      = "Code"
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

		if !utils.IsGlobalAdmin(ctx) {
			l.Warn("cannot access country page without global admin role")
			http.Error(w, "user is not global admin", http.StatusForbidden)
			return
		}

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
		country.Name = r.FormValue(formParamName)
		country.Code = r.FormValue(formParamCode)

		country, err = repo.Put(r.Context(), country)
		if err != nil {
			l.Error("failed to put country", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = client.AddCountry(r.Context(), country.Code)

		if err != nil {
			l.Error("failed to add country to zanzibar graph", zap.Error(err))
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
