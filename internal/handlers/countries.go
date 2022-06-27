package handlers

import (
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func HandleCountries(templates map[string]*template.Template, repo db.CountryRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const templateName = "countries.gohtml"

		var (
			ctx       = r.Context()
			l         = logging.NewLogger(ctx)
			countries []*api.Country
			err       error
		)

		if countries, err = repo.GetAll(ctx); err != nil {
			l.Error("failed to get countries", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		renderView(templates, templateName, w, r, map[string]interface{}{
			"Countries": countries,
		})
	})
}
