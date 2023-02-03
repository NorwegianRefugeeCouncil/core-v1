package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleHome(templates map[string]*template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		authInterface, err := utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "couldn't get auth context: "+err.Error(), http.StatusInternalServerError)
			return
		}

		allowedCountryIDs := authInterface.GetAllowedCountries().Items()

		if authInterface.IsGlobalAdmin() || len(allowedCountryIDs) != 1 {
			http.Redirect(w, r, "/countries", http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, fmt.Sprintf("/countries/%s/participants", allowedCountryIDs[0]), http.StatusTemporaryRedirect)
		}
	})
}
