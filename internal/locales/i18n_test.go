package locales

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBColumns(t *testing.T) {
	LoadTranslations()
	Init()

	testCases := []struct {
		name        string
		values      []string
		expect      []string
		expectError error
	}{
		{
			name:        "empty",
			values:      []string{},
			expect:      []string{},
			expectError: nil,
		},
		{
			name:        "default language",
			values:      []string{"ID", "First name"},
			expect:      []string{"id", "first_name"},
			expectError: nil,
		},
		{
			name:        "debug language",
			values:      []string{"XXXX_file_id", "XXXX_file_first_name"},
			expect:      []string{"id", "first_name"},
			expectError: nil,
		},
		{
			name:        "db columns",
			values:      []string{"id", "first_name"},
			expect:      []string{"id", "first_name"},
			expectError: nil,
		},
		{
			name:        "mixing languages",
			values:      []string{"XXXX_file_id", "first_name", "Full name"},
			expect:      []string{"id", "first_name", "full_name"},
			expectError: nil,
		},
		{
			name:        "unknown columns",
			values:      []string{"XXXX_file_id", "first_name", "Full name", "Other", "Unknown", ""},
			expect:      []string{},
			expectError: fmt.Errorf(l.Translate("error_unknown_columns", "Other, Unknown, <empty>")),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			keys, errs := GetDBColumns(tt.values)
			if tt.expectError != nil {
				assert.EqualValues(t, tt.expectError, errs)
				return
			}

			assert.Equal(t, tt.expect, keys)
		})
	}

}
