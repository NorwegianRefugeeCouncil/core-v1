package middleware

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nrc-no/notcore/internal/locales"
)

func Localize(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		acceptLang := r.Header.Get("Accept-Language")
		if q := r.URL.Query().Get("lang"); q != "" {
			acceptLang = q
		}
		if !locales.AvailableLangs.Contains(acceptLang) {
			acceptLang = locales.DefaultLang.String()
		}
		localizer := i18n.NewLocalizer(locales.Translations, acceptLang)
		ctx = locales.WithLocalizer(ctx, localizer)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
