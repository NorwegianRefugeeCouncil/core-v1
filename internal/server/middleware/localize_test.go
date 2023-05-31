package middleware

import (
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_getAppropriateLanguage(t *testing.T) {
	tests := []struct {
		name     string
		header   string
		urlParam http.Cookie
		result   string
	}{
		{
			name:     "header is empty, url param is empty, not available",
			header:   "",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: ""},
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is empty, url param is not empty, not available",
			header:   "",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: "el"},
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is not empty, url param is empty, not available",
			header:   "el",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: ""},
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is not empty, url param is not empty, not available",
			header:   "el",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: "jp"},
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is empty, url param is not empty, available",
			header:   "",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: "es"},
			result:   "es",
		},
		{
			name:     "header is not empty, url param is empty, available",
			header:   "es",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: ""},
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, only header available",
			header:   "es",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: "jp"},
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, only url available",
			header:   "jp",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: "es"},
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, both available",
			header:   "en",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: "es"},
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, both available",
			header:   "es",
			urlParam: http.Cookie{Name: "nrc-core-language", Value: "en"},
			result:   "en",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			language := getAppropriateLanguage(tt.header, &tt.urlParam, containers.NewStringSet("en", "es"))
			assert.Equal(t, tt.result, language)
		})
	}
}
