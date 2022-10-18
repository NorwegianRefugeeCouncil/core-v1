package handlers

import (
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleCountrySelector(templates map[string]*template.Template) http.Handler {

	const (
		templateName               = "country_select.gohtml"
		viewParamsAllowedCountries = "AllowedCountries"
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

		allowedCountryIDs := authCtx.GetCountryIDsWithPermission("read")
		countries, err := utils.GetCountries(ctx)
		if err != nil {
			l.Error("failed to get countries", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var allowedCountries []*api.Country
		for _, c := range countries {
			if allowedCountryIDs.Contains(c.ID) {
				allowedCountries = append(allowedCountries, c)
			}
		}

		renderView(templates, templateName, w, r, viewParams{
			viewParamsAllowedCountries: allowedCountries,
		})
	})
}
