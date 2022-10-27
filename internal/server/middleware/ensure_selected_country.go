package middleware

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func EnsureSelectedCountry() func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			redirect := func() {
				http.Redirect(w, r, "/countries", http.StatusTemporaryRedirect)
			}

			selectedCountryID, err := utils.GetSelectedCountryID(ctx)
			if err != nil {
				l.Error("failed to get selected country id", zap.Error(err))
				redirect()
				return
			}

			if len(selectedCountryID) == 0 {
				l.Debug("no selected country. redirecting")
				redirect()
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
