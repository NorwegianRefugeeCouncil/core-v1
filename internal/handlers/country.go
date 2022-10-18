package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleCountry(templates map[string]*template.Template, repo db.CountryRepo) http.Handler {

	const (
		templateName       = "country.gohtml"
		newId              = "new"
		pathParamCountryID = "country_id"
		viewParamCountry   = "Country"
		formParamName      = "Name"
		formParamCode      = "Code"
		formParamJwtGroup  = "JwtGroup"
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

		authCtx, err := utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !authCtx.IsGlobalAdmin() {
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
		country.Name = strings.TrimSpace(r.FormValue(formParamName))
		country.Code = strings.TrimSpace(strings.ToLower(r.FormValue(formParamCode)))
		country.JwtGroup = strings.TrimSpace(r.FormValue(formParamJwtGroup))

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
