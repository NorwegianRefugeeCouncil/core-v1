package handlers

import (
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleCountries(templates map[string]*template.Template, countryRepo db.CountryRepo) http.Handler {

	const (
		templateName = "countries.gohtml"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
			err error
		)

		authCtx, err := utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !authCtx.IsGlobalAdmin() {
			l.Warn("User is not global admin")
			http.Error(w, "You are not allowed to access this page", http.StatusForbidden)
			return
		}

		render := func() {
			renderView(templates, templateName, w, r, viewParams{})
		}

		render()
	})
}
