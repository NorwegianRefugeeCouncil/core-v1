package middleware

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/locales"
	"net/http"
	"strings"
)

const cookieName = "nrc-core-language"

func Localize(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// TODO: activate when translations are available
		//languageCookie, _ := r.Cookie(cookieName)
		//languageHeader := r.Header.Get("Accept-Language")
		//language := getAppropriateLanguage(languageHeader, languageCookie, locales.AvailableLangs)
		localizer := i18n.NewLocalizer(locales.Translations, locales.DefaultLang.String())
		ctx = locales.WithLocalizer(ctx, localizer)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func getAppropriateLanguage(languageHeader string, cookie *http.Cookie, availableLangs containers.StringSet) string {
	if cookie != nil && availableLangs.Contains(cookie.Value) {
		return cookie.Value
	}

	headerLanguages := []string{}
	for _, lang := range strings.Split(languageHeader, ",") {
		headerLanguages = append(headerLanguages, strings.Split(lang, ";")[0])
	}
	for _, l := range headerLanguages {
		if availableLangs.Contains(l) {
			return l
		}
	}

	return locales.DefaultLang.String()
}
