package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nrc-no/notcore/internal/containers"
)

func TestHasGlobalAdminPermission(t *testing.T) {
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
				HasGlobalAdminPermission()(
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
