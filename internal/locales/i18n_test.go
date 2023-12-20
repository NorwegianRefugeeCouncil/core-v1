package locales

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTranslationKeys(t *testing.T) {
	LoadTranslations()
	Init()

	testCases := []struct {
		name   string
		values []string
		expect []string
	}{
		{
			name:   "empty",
			values: []string{},
			expect: []string{},
		},
		{
			name:   "default language",
			values: []string{"id", "first_name"},
			expect: []string{"file_id", "file_first_name"},
		},
		{
			name:   "debug language",
			values: []string{"XXXX_file_id", "XXXX_file_first_name"},
			expect: []string{"file_id", "file_first_name"},
		},
		{
			name:   "mixing languages",
			values: []string{"XXXX_file_id", "first_name"},
			expect: []string{"file_id", "file_first_name"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			keys := GetTranslationKeys(tt.values)
			assert.Equal(t, tt.expect, keys)
		})
	}

}
