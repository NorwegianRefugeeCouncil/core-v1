package middleware

import (
	"github.com/nrc-no/notcore/internal/containers"
	"net/http"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nrc-no/notcore/internal/locales"
)

func Localize(next http.Handler) http.Handler {
	localizationHandler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		language := getAppropriateLanguage(r.Header.Get("Accept-Language"), r.URL.Query().Get("lang"), locales.AvailableLangs)
		localizer := i18n.NewLocalizer(locales.Translations, language)
		ctx = locales.WithLocalizer(ctx, localizer)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(localizationHandler)
}

func getAppropriateLanguage(languageHeader string, urlParam string, availableLangs containers.StringSet) string {
	language := locales.DefaultLang.String()

	headerLanguages := []string{}
	for _, lang := range strings.Split(languageHeader, ",") {
		headerLanguages = append(headerLanguages, strings.Split(lang, ";")[0])
	}
	for _, l := range headerLanguages {
		if availableLangs.Contains(l) {
			language = l
			break
		}
	}

	if availableLangs.Contains(urlParam) {
		language = urlParam
	}
	return language
}
