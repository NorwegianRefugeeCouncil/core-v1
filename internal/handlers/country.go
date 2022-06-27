package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func HandleCountry(templates map[string]*template.Template, repo db.CountryRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const templateName = "country.gohtml"

		var (
			ctx       = r.Context()
			l         = logging.NewLogger(ctx)
			err       error
			country   = &api.Country{}
			countryID = mux.Vars(r)["country_id"]
		)

		render := func() {
			renderView(templates, templateName, w, r, map[string]interface{}{
				"Country": country,
			})
		}

		if r.Method == "GET" {
			if countryID != "new" {
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

		isNew := countryID == "new"
		if !isNew {
			country.ID = countryID
		}
		country.Name = r.FormValue("Name")
		country.Code = r.FormValue("Code")

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
