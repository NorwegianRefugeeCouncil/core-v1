package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/utils"
)

func HandleIndividual(templates map[string]*template.Template, repo db.IndividualRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var individual = &api.Individual{}

		individualId := mux.Vars(r)["individual_id"]

		render := func() {
			if individual == nil {
				individual = &api.Individual{}
			}
			if err := templates["individual.gohtml"].ExecuteTemplate(w, "base", map[string]interface{}{
				"Individual": individual,
			}); err != nil {
				println(err.Error())
			}
			return
		}

		individualId = mux.Vars(r)["individual_id"]
		if individualId != "new" {
			individual, err = repo.GetByID(r.Context(), individualId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if r.Method == "GET" {
			render()
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		individual.FullName = r.FormValue("FullName")
		individual.Email = utils.NormalizeEmail(r.FormValue("Email"))
		individual.PhoneNumber = r.FormValue("PhoneNumber")
		individual.Address = r.FormValue("Address")
		individual.Gender = r.FormValue("Gender")
		individual.NormalizedPhoneNumber = utils.NormalizePhoneNumber(individual.PhoneNumber)
		birthDateStr := r.FormValue("BirthDate")
		if birthDateStr != "" {
			birthDate, err := time.Parse("2006-01-02", birthDateStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			individual.BirthDate = &birthDate
		}
		_, err = repo.Put(r.Context(), individual)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/individuals/"+individual.ID, http.StatusSeeOther)

	})
}
