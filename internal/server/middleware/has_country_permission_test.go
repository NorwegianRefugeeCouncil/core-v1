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

func TestHasCountryPermission(t *testing.T) {
	var parametrizedTestsCanRead = []struct {
		name              string
		allowedCountryIDs containers.StringSet
		allCountryIDs     containers.StringSet
		isGlobalAdmin     bool
		canRead           bool
		canWrite          bool
		selectedCountryID string
		expectedStatus    int
	}{
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1"), false, true, false, "1", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] canRead; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1", "2"), false, true, false, "1", http.StatusOK},
		{"[Allowed] two;   [All] two;   [PermLevel] canRead; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2"), false, true, false, "2", http.StatusOK},
		{"[Allowed] two;   [All] three; [PermLevel] canRead; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2", "3"), false, true, false, "2", http.StatusOK},
		{"[Allowed] empty; [All] empty; [PermLevel] canRead; [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), false, true, false, "0", http.StatusForbidden},
		{"[Allowed] empty; [All] empty; [PermLevel] canRead; [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), false, true, false, "0", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In allowed only", containers.NewStringSet("1"), containers.NewStringSet("2"), false, true, false, "1", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2"), false, true, false, "2", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canRead; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("2"), false, true, false, "1", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, true, false, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, true, false, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, true, false, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canRead; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), false, true, false, "3", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canRead; [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, true, false, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canRead; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), false, true, false, "3", http.StatusForbidden},
	}

	var parametrizedTestsCanWrite = []struct {
		name              string
		allowedCountryIDs containers.StringSet
		allCountryIDs     containers.StringSet
		isGlobalAdmin     bool
		canRead           bool
		canWrite          bool
		selectedCountryID string
		expectedStatus    int
	}{
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1"), false, false, true, "1", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] canWrite; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1", "2"), false, false, true, "1", http.StatusOK},
		{"[Allowed] two;   [All] two;   [PermLevel] canWrite; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2"), false, false, true, "2", http.StatusOK},
		{"[Allowed] two;   [All] three; [PermLevel] canWrite; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2", "3"), false, false, true, "2", http.StatusOK},
		{"[Allowed] empty; [All] empty; [PermLevel] canWrite; [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), false, false, true, "0", http.StatusForbidden},
		{"[Allowed] empty; [All] empty; [PermLevel] canWrite; [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), false, false, true, "0", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In allowed only", containers.NewStringSet("1"), containers.NewStringSet("2"), false, false, true, "1", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2"), false, false, true, "2", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [PermLevel] canWrite; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("2"), false, false, true, "1", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, false, true, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, false, true, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, false, true, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canWrite; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), false, false, true, "3", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [PermLevel] canWrite; [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, false, true, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [PermLevel] canWrite; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), false, false, true, "3", http.StatusForbidden},
	}

	var parametrizedTestsIsGlobalAdmin = []struct {
		name              string
		allowedCountryIDs containers.StringSet
		allCountryIDs     containers.StringSet
		isGlobalAdmin     bool
		canRead           bool
		canWrite          bool
		selectedCountryID string
		expectedStatus    int
	}{
		{"[Allowed] empty; [All] empty; [PermLevel] GlobalAdmin; [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), true, false, false, "0", http.StatusOK},
		{"[Allowed] empty; [All] empty; [PermLevel] GlobalAdmin; [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), true, false, false, "0", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1"), true, false, false, "1", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In allowed only", containers.NewStringSet("1"), containers.NewStringSet("2"), true, false, false, "1", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2"), true, false, false, "2", http.StatusOK},
		{"[Allowed] one;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("2"), true, false, false, "1", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1", "2"), true, false, false, "1", http.StatusOK},
		{"[Allowed] two;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2"), true, false, false, "2", http.StatusOK},
		{"[Allowed] two;   [All] three; [PermLevel] GlobalAdmin; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2", "3"), true, false, false, "2", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, false, false, "2", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, false, false, "2", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, false, false, "3", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), true, false, false, "3", http.StatusOK},
		{"[Allowed] two;   [All] one;   [PermLevel] GlobalAdmin; [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, false, false, "3", http.StatusOK},
		{"[Allowed] one;   [All] two;   [PermLevel] GlobalAdmin; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), true, false, false, "3", http.StatusOK},
	}

	for _, tt := range parametrizedTestsCanRead {
		t.Run(fmt.Sprintf("[Permission] %s; %s", string(auth.PermissionRead), tt.name), func(t *testing.T) {
			handlerToTest := configureDummyContextMiddleware(
				tt.allowedCountryIDs,
				tt.allCountryIDs,
				tt.isGlobalAdmin,
				tt.canRead,
				tt.canWrite,
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
				tt.allowedCountryIDs,
				tt.allCountryIDs,
				tt.isGlobalAdmin,
				tt.canRead,
				tt.canWrite,
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
				tt.allowedCountryIDs,
				tt.allCountryIDs,
				tt.isGlobalAdmin,
				tt.canRead,
				tt.canWrite,
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
