package middleware

import (
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_getAppropriateLanguage(t *testing.T) {
	locales.LoadTranslations()
	locales.Init()
	tests := []struct {
		name   string
		header string
		cookie http.Cookie
		result string
	}{
		{
			name:   "header is empty, url param is empty, not available",
			header: "",
			cookie: http.Cookie{Name: "nrc-core-language", Value: ""},
			result: locales.DefaultLang.String(),
		},
		{
			name:   "header is empty, url param is not empty, not available",
			header: "",
			cookie: http.Cookie{Name: "nrc-core-language", Value: "el"},
			result: locales.DefaultLang.String(),
		},
		{
			name:   "header is not empty, url param is empty, not available",
			header: "el",
			cookie: http.Cookie{Name: "nrc-core-language", Value: ""},
			result: locales.DefaultLang.String(),
		},
		{
			name:   "header is not empty, url param is not empty, not available",
			header: "el",
			cookie: http.Cookie{Name: "nrc-core-language", Value: "jp"},
			result: locales.DefaultLang.String(),
		},
		{
			name:   "header is empty, url param is not empty, available",
			header: "",
			cookie: http.Cookie{Name: "nrc-core-language", Value: "ja"},
			result: "ja",
		},
		{
			name:   "header is not empty, url param is empty, available",
			header: "ja",
			cookie: http.Cookie{Name: "nrc-core-language", Value: ""},
			result: "ja",
		},
		{
			name:   "header is not empty, url param is not empty, only header available",
			header: "ja",
			cookie: http.Cookie{Name: "nrc-core-language", Value: "jp"},
			result: "ja",
		},
		{
			name:   "header is not empty, url param is not empty, only url available",
			header: "jp",
			cookie: http.Cookie{Name: "nrc-core-language", Value: "ja"},
			result: "ja",
		},
		{
			name:   "header is not empty, url param is not empty, both available",
			header: "en",
			cookie: http.Cookie{Name: "nrc-core-language", Value: "ja"},
			result: "ja",
		},
		{
			name:   "header is not empty, url param is not empty, both available",
			header: "ja",
			cookie: http.Cookie{Name: "nrc-core-language", Value: "en"},
			result: "en",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			language := getAppropriateLanguage(tt.header, &tt.cookie)
			assert.Equal(t, tt.result, language)
		})
	}
}
