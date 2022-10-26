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
	var permissions = []struct {
		permission auth.Permission
	}{
		{permission: auth.PermissionRead},
		{permission: auth.PermissionWrite},
	}

	var parametrizedTests = []struct {
		name              string
		allowedCountryIDs containers.Set[string]
		allCountryIDs     containers.Set[string]
		isGlobalAdmin     bool
		selectedCountryID string
		expectedStatus    int
	}{
		{"[Allowed] empty; [All] empty; [GlobalAdmin] true;  [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), true, "0", http.StatusForbidden},
		{"[Allowed] empty; [All] empty; [GlobalAdmin] false; [Selected] In neither", containers.NewStringSet(), containers.NewStringSet(), false, "0", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [GlobalAdmin] false; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1"), false, "1", http.StatusOK},
		{"[Allowed] one;   [All] one;   [GlobalAdmin] false; [Selected] In allowed only", containers.NewStringSet("1"), containers.NewStringSet("2"), false, "1", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [GlobalAdmin] false; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2"), false, "2", http.StatusForbidden},
		{"[Allowed] one;   [All] one;   [GlobalAdmin] true;  [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("2"), true, "1", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [GlobalAdmin] false; [Selected] In both", containers.NewStringSet("1"), containers.NewStringSet("1", "2"), false, "1", http.StatusOK},
		{"[Allowed] two;   [All] two;   [GlobalAdmin] false; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2"), false, "2", http.StatusOK},
		{"[Allowed] two;   [All] three; [GlobalAdmin] false; [Selected] In both", containers.NewStringSet("1", "2"), containers.NewStringSet("1", "2", "3"), false, "2", http.StatusOK},
		{"[Allowed] two;   [All] one;   [GlobalAdmin] false; [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [GlobalAdmin] true;  [Selected] In allowed only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, "2", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [GlobalAdmin] false; [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), false, "3", http.StatusForbidden},
		{"[Allowed] one;   [All] two;   [GlobalAdmin] false; [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), false, "3", http.StatusForbidden},
		{"[Allowed] two;   [All] one;   [GlobalAdmin] true;  [Selected] In all only", containers.NewStringSet("1", "2"), containers.NewStringSet("3"), true, "3", http.StatusOK},
		{"[Allowed] one;   [All] two;   [GlobalAdmin] true;  [Selected] In all only", containers.NewStringSet("1"), containers.NewStringSet("2", "3"), true, "3", http.StatusOK},
	}

	for _, pp := range permissions {
		for _, tt := range parametrizedTests {
			t.Run(fmt.Sprintf("[Permission] %s; %s", string(pp.permission), tt.name), func(t *testing.T) {
				handlerToTest := configureDummyContextMiddleware(
					tt.allowedCountryIDs,
					tt.allCountryIDs,
					tt.isGlobalAdmin,
					tt.selectedCountryID,
				)(
					HasCountryPermission(pp.permission)(
						nextHandler(),
					),
				)
				req := httptest.NewRequest("GET", "http://testing", nil)
				responeRecorder := httptest.NewRecorder()
				handlerToTest.ServeHTTP(responeRecorder, req)
				assert.Equal(t, tt.expectedStatus, responeRecorder.Code)
			})
		}
	}
}
