package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils"
)

func TestHasGlobalAdminPermissionMiddleware(t *testing.T) {
	var parametrizedTests = []struct {
		isGlobalAdmin  bool
		expectedStatus int
	}{
		{true, http.StatusOK},
		{false, http.StatusForbidden},
	}

	for _, tt := range parametrizedTests {
		t.Run("", func(t *testing.T) {
			handlerToTest := configureDummyContextMiddleware(
				containers.NewStringSet(),
				containers.NewStringSet(),
				tt.isGlobalAdmin,
				"",
			)(
				hasGlobalAdminPermissionMiddleware()(
					nextHandler(),
				),
			)
			req := httptest.NewRequest("GET", "http://testing", nil)
			responeRecorder := httptest.NewRecorder()
			handlerToTest.ServeHTTP(responeRecorder, req)
			if responeRecorder.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, responeRecorder.Code)
			}
		})
	}
}

func nextHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func configureDummyContextMiddleware(
	allowedCountryIDs containers.Set[string],
	allCountryIDs containers.Set[string],
	isGlobalAdmin bool,
	selectedCountryId string,
) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authIntf := auth.New(
				allowedCountryIDs,
				allCountryIDs,
				isGlobalAdmin,
			)
			r = r.WithContext(
				utils.WithAuthContext(
					utils.WithSelectedCountryID(r.Context(), selectedCountryId),
					authIntf,
				),
			)
			next.ServeHTTP(w, r)
		})
	}
}
