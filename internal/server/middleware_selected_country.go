package server

import (
	"net/http"
	"time"

	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func selectedCountryMiddleware() func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		const cookieName = "core_selectedCountryID"

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			l.Debug("selected country middleware")
			selectedCountryIDCookie, err := r.Cookie(cookieName)
			if err != nil {
				l.Debug("selected country cookie not found")
				h.ServeHTTP(w, r)
				return
			}

			l.Debug("selected country cookie value", zap.String("cookie", selectedCountryIDCookie.Value))
			allCountries, err := utils.GetCountries(ctx)
			if err != nil {
				l.Error("failed to get all countries", zap.Error(err))
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			// checking if the selected country id is actually a valid country id
			found := false
			for _, c := range allCountries {
				if c.ID == selectedCountryIDCookie.Value {
					found = true
					break
				}
			}

			if found {
				// Store the selected country ID in the request context
				r = r.WithContext(utils.WithSelectedCountryID(ctx, selectedCountryIDCookie.Value))
			} else {
				// Remove the cookie
				http.SetCookie(w, &http.Cookie{
					Name:     cookieName,
					Value:    "",
					Path:     "/",
					Expires:  time.Unix(0, 0),
					HttpOnly: true,
					Secure:   true,
				})
			}
			h.ServeHTTP(w, r)
		})
	}
}
