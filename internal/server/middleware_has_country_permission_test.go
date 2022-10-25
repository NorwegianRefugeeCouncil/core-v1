package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/stretchr/testify/assert"
)

func TestHasCountryPermissionMiddleware(t *testing.T) {
	var methodParametrizedTests = []struct {
		method  string
		allowed bool
	}{
		{"GET", true},
		{"POST", true},
		{"PUT", true},
		{"DELETE", true},
		{"PATCH", false},
	}

	var parametrizedTests = []struct {
		allowedCountryIDs containers.Set[string]
		allCountryIDs     containers.Set[string]
		isGlobalAdmin     bool
		selectedCountryID string
		expectedStatus    int
	}{
		{containers.NewStringSet(), containers.NewStringSet(), true, "0", http.StatusForbidden},
		{containers.NewStringSet(), containers.NewStringSet(), false, "0", http.StatusForbidden},
		{containers.NewStringSet("1"), containers.NewStringSet("1"), false, "1", http.StatusOK},
		{containers.NewStringSet("1"), containers.NewStringSet("2"), false, "1", http.StatusForbidden},
		{containers.NewStringSet("1"), containers.NewStringSet("2"), false, "2", http.StatusForbidden},
		{containers.NewStringSet("1"), containers.NewStringSet("2"), true, "1", http.StatusForbidden},
		{containers.NewStringSet("1"), containers.NewStringSet("1", "2"), false, "1", http.StatusOK},
		{containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2"), false, "2", http.StatusOK},
		{containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2", "3"), false, "2", http.StatusOK},
		{containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, "2", http.StatusForbidden},
		{containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, "2", http.StatusForbidden},
		{containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, "3", http.StatusForbidden},
		{containers.NewStringSet("1"), containers.NewStringSet("2", "3"), false, "3", http.StatusForbidden},
		{containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, "3", http.StatusOK},
		{containers.NewStringSet("1"), containers.NewStringSet("2", "3"), true, "3", http.StatusOK},
	}

	for _, mm := range methodParametrizedTests {
		for _, tt := range parametrizedTests {
			t.Run("", func(t *testing.T) {
				handlerToTest := configureDummyContextMiddleware(
					tt.allowedCountryIDs,
					tt.allCountryIDs,
					tt.isGlobalAdmin,
					tt.selectedCountryID,
				)(
					hasCountryPermissionMiddleware()(
						nextHandler(),
					),
				)
				req := httptest.NewRequest(mm.method, "http://testing", nil)
				responeRecorder := httptest.NewRecorder()
				handlerToTest.ServeHTTP(responeRecorder, req)
				if mm.allowed {
					assert.Equal(t, tt.expectedStatus, responeRecorder.Code)
				} else {
					assert.Equal(t, http.StatusMethodNotAllowed, responeRecorder.Code)
				}
			})
		}
	}
}
