package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
)

func HandleCountry(templates map[string]*template.Template, repo db.CountryRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var country = &api.Country{}
		countryID := mux.Vars(r)["country_id"]

		if r.Method == "GET" {

			if countryID != "new" {
				country, err = repo.GetByID(r.Context(), countryID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			templates["country.gohtml"].ExecuteTemplate(w, "base", map[string]interface{}{
				"Country": country,
			})
			return
		}

		if countryID != "new" {
			country.ID = countryID
		}
		country.Name = r.FormValue("Name")
		country.Code = r.FormValue("Code")

		ret, err := repo.Put(r.Context(), country)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/countries/"+ret.ID, http.StatusSeeOther)

	})
}
