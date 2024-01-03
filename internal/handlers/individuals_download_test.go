package handlers

import (
	"testing"
)

func TestParseRangeHeader(t *testing.T) {
	tests := []struct {
		name                 string
		rangeHeader          string
		expectedOffset       int64
		expectedCount        int64
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:           "valid range header",
			rangeHeader:    "bytes=100-200",
			expectedOffset: 100,
			expectedCount:  200,
			expectedError:  false,
		},
		{
			name:                 "invalid range header format",
			rangeHeader:          "invalid",
			expectedError:        true,
			expectedErrorMessage: "invalid range header",
		},
		{
			name:                 "invalid offset",
			rangeHeader:          "bytes=abc-200",
			expectedError:        true,
			expectedErrorMessage: "invalid range header",
		},
		{
			name:                 "invalid count",
			rangeHeader:          "bytes=100-xyz",
			expectedError:        true,
			expectedErrorMessage: "invalid range header",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			offset, count, err := parseRangeHeader(test.rangeHeader)

			if test.expectedError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				} else if err.Error() != test.expectedErrorMessage {
					t.Errorf("Expected error message '%s', but got '%s'", test.expectedErrorMessage, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got '%s'", err.Error())
				} else {
					if offset != test.expectedOffset {
						t.Errorf("Expected offset %d, but got %d", test.expectedOffset, offset)
					}
					if count != test.expectedCount {
						t.Errorf("Expected count %d, but got %d", test.expectedCount, count)
					}
				}
			}
		})
	}
}
