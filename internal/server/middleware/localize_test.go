package middleware

import (
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getAppropriateLanguage(t *testing.T) {
	tests := []struct {
		name     string
		header   string
		urlParam string
		result   string
	}{
		{
			name:     "header is empty, url param is empty, not available",
			header:   "",
			urlParam: "",
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is empty, url param is not empty, not available",
			header:   "",
			urlParam: "el",
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is not empty, url param is empty, not available",
			header:   "el",
			urlParam: "",
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is not empty, url param is not empty, not available",
			header:   "el",
			urlParam: "jp",
			result:   locales.DefaultLang.String(),
		},
		{
			name:     "header is empty, url param is not empty, available",
			header:   "",
			urlParam: "es",
			result:   "es",
		},
		{
			name:     "header is not empty, url param is empty, available",
			header:   "es",
			urlParam: "",
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, only header available",
			header:   "es",
			urlParam: "jp",
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, only url available",
			header:   "jp",
			urlParam: "es",
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, both available",
			header:   "en",
			urlParam: "es",
			result:   "es",
		},
		{
			name:     "header is not empty, url param is not empty, both available",
			header:   "es",
			urlParam: "en",
			result:   "en",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			language := getAppropriateLanguage(tt.header, tt.urlParam, containers.NewStringSet("en", "es"))
			assert.Equal(t, tt.result, language)
		})
	}
}
