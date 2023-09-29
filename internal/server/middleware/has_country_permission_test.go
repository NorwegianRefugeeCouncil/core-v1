package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/stretchr/testify/assert"
)

func makeReadCountryPermissions(countryIDs ...string) auth.CountryPermissions {
	perms := make(auth.CountryPermissions)
	for _, id := range countryIDs {
		perms[id] = containers.NewSet(auth.PermissionRead)
	}
	return perms
}

func makeWriteCountryPermissions(countryIDs ...string) auth.CountryPermissions {
	perms := make(auth.CountryPermissions)
	for _, id := range countryIDs {
		perms[id] = containers.NewSet(auth.PermissionWrite)
	}
	return perms
}

func TestHasCountryPermission(t *testing.T) {
	var parametrizedTestsCanRead = []struct {
		name              string
		countryPermissions auth.CountryPermissions
		allCountryIDs     containers.StringSet
		isGlobalAdmin     bool
		selectedCountryID string
		expectedStatus    int
	}{
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In both", makeReadCountryPermissions("1"), containers.NewStringSet("1"), false, "1", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] canRead; [Selected] In both", makeReadCountryPermissions("1"), containers.NewStringSet("1", "2"), false, "1", http.StatusOK},
		{"[Allowed] two;   [All] two;   [PermLevel] canRead; [Selected] In both", makeReadCountryPermissions("1", "2"), containers.NewStringSet("1", "2"), false, "2", http.StatusOK},
		{"[Allowed] two;   [All] three; [PermLevel] canRead; [Selected] In both", makeReadCountryPermissions("1", "2"), containers.NewStringSet("1", "2", "3"), false, "2", http.StatusOK},
		{"[Allowed] empty; [All] empty; [PermLevel] canRead; [Selected] In neither", makeReadCountryPermissions(), containers.NewStringSet(), false, "0", http.StatusForbidden},
		{"[Allowed] empty; [All] empty; [PermLevel] canRead; [Selected] In neither", makeReadCountryPermissions(), containers.NewStringSet(), false, "0", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In allowed only", makeReadCountryPermissions("1"), containers.NewStringSet("2"), false, "1", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In all only", makeReadCountryPermissions("1"), containers.NewStringSet("2"), false, "2", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In both", makeReadCountryPermissions("1"), containers.NewStringSet("2"), false, "1", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In allowed only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In allowed only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In all only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canRead; [Selected] In all only", makeReadCountryPermissions("1"), containers.NewStringSet("2", "3"), false, "3", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In all only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canRead; [Selected] In all only", makeReadCountryPermissions("1"), containers.NewStringSet("2", "3"), false, "3", http.StatusForbidden},
	}

	var parametrizedTestsCanWrite = []struct {
		name              string
		countryPermissions auth.CountryPermissions
		allCountryIDs     containers.StringSet
		isGlobalAdmin     bool
		selectedCountryID string
		expectedStatus    int
	}{
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In both", makeWriteCountryPermissions("1"), containers.NewStringSet("1"), false, "1", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] canWrite; [Selected] In both", makeWriteCountryPermissions("1"), containers.NewStringSet("1", "2"), false, "1", http.StatusOK},
		{"[Allowed] two;   [All] two;   [PermLevel] canWrite; [Selected] In both", makeWriteCountryPermissions("1", "2"), containers.NewStringSet("1", "2"), false, "2", http.StatusOK},
		{"[Allowed] two;   [All] three; [PermLevel] canWrite; [Selected] In both", makeWriteCountryPermissions("1", "2"), containers.NewStringSet("1", "2", "3"), false, "2", http.StatusOK},
		{"[Allowed] empty; [All] empty; [PermLevel] canWrite; [Selected] In neither", makeWriteCountryPermissions(), containers.NewStringSet(), false, "0", http.StatusForbidden},
		{"[Allowed] empty; [All] empty; [PermLevel] canWrite; [Selected] In neither", makeWriteCountryPermissions(), containers.NewStringSet(), false, "0", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In allowed only", makeWriteCountryPermissions("1"), containers.NewStringSet("2"), false, "1", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In all only", makeWriteCountryPermissions("1"), containers.NewStringSet("2"), false, "2", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In both", makeWriteCountryPermissions("1"), containers.NewStringSet("2"), false, "1", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In allowed only", makeWriteCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In allowed only", makeWriteCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In all only", makeWriteCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canWrite; [Selected] In all only", makeWriteCountryPermissions("1"), containers.NewStringSet("2", "3"), false, "3", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In all only", makeWriteCountryPermissions("1", "2"), containers.NewStringSet("3"), false, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canWrite; [Selected] In all only", makeWriteCountryPermissions("1"), containers.NewStringSet("2", "3"), false, "3", http.StatusForbidden},
	}

	var parametrizedTestsIsGlobalAdmin = []struct {
		name              string
		countryPermissions auth.CountryPermissions
		allCountryIDs     containers.StringSet
		isGlobalAdmin     bool
		selectedCountryID string
		expectedStatus    int
	}{
		{"[Allowed] empty; [All] empty; [PermLevel] GlobalAdmin; [Selected] In neither", makeReadCountryPermissions(), containers.NewStringSet(), true, "0", http.StatusOK},
		{"[Allowed] empty; [All] empty; [PermLevel] GlobalAdmin; [Selected] In neither", makeReadCountryPermissions(), containers.NewStringSet(), true, "0", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In both", makeReadCountryPermissions("1"), containers.NewStringSet("1"), true, "1", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In allowed only", makeReadCountryPermissions("1"), containers.NewStringSet("2"), true, "1", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In all only", makeReadCountryPermissions("1"), containers.NewStringSet("2"), true, "2", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In both", makeReadCountryPermissions("1"), containers.NewStringSet("2"), true, "1", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In both", makeReadCountryPermissions("1"), containers.NewStringSet("1", "2"), true, "1", http.StatusOK},
		{"[Allowed] two;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In both", makeReadCountryPermissions("1", "2"), containers.NewStringSet("1", "2"), true, "2", http.StatusOK},
		{"[Allowed] two;   [All] three; [PermLevel] GlobalAdmin; [Selected] In both", makeReadCountryPermissions("1", "2"), containers.NewStringSet("1", "2", "3"), true, "2", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In allowed only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), true, "2", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In allowed only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), true, "2", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In all only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), true, "3", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In all only", makeReadCountryPermissions("1"), containers.NewStringSet("2", "3"), true, "3", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In all only", makeReadCountryPermissions("1", "2"), containers.NewStringSet("3"), true, "3", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In all only", makeReadCountryPermissions("1"), containers.NewStringSet("2", "3"), true, "3", http.StatusOK},
	}

	for _, tt := range parametrizedTestsCanRead {
		t.Run(fmt.Sprintf("[Permission] %s; %s", string(auth.PermissionRead), tt.name), func(t *testing.T) {
			handlerToTest := configureDummyContextMiddleware(
				tt.countryPermissions,
				tt.allCountryIDs,
				tt.isGlobalAdmin,
				tt.selectedCountryID,
			)(
				HasCountryPermission(auth.PermissionRead)(
					nextHandler(),
				),
			)
			req := httptest.NewRequest("GET", "http://testing", nil)
			responseRecorder := httptest.NewRecorder()
			handlerToTest.ServeHTTP(responseRecorder, req)
			assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
		})
	}

	for _, tt := range parametrizedTestsCanWrite {
		t.Run(fmt.Sprintf("[Permission] %s; %s", string(auth.PermissionWrite), tt.name), func(t *testing.T) {
			handlerToTest := configureDummyContextMiddleware(
				tt.countryPermissions,
				tt.allCountryIDs,
				tt.isGlobalAdmin,
				tt.selectedCountryID,
			)(
				HasCountryPermission(auth.PermissionWrite)(
					nextHandler(),
				),
			)
			req := httptest.NewRequest("GET", "http://testing", nil)
			responseRecorder := httptest.NewRecorder()
			handlerToTest.ServeHTTP(responseRecorder, req)
			assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
		})
	}

	for _, tt := range parametrizedTestsIsGlobalAdmin {
		t.Run(fmt.Sprintf("[Permission] %s; %s", string(auth.PermissionGlobalAdmin), tt.name), func(t *testing.T) {
			handlerToTest := configureDummyContextMiddleware(
				tt.countryPermissions,
				tt.allCountryIDs,
				tt.isGlobalAdmin,
				tt.selectedCountryID,
			)(
				HasCountryPermission(auth.PermissionGlobalAdmin)(
					nextHandler(),
				),
			)
			req := httptest.NewRequest("GET", "http://testing", nil)
			responseRecorder := httptest.NewRecorder()
			handlerToTest.ServeHTTP(responseRecorder, req)
			assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
		})
	}

}
