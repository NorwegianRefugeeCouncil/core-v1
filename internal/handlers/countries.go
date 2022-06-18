package handlers

import (
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleCountries(templates map[string]*template.Template, repo db.CountryRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		countries, err := repo.GetAll(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := templates["countries.gohtml"].ExecuteTemplate(w, "base", map[string]interface{}{
			"Countries": countries,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
