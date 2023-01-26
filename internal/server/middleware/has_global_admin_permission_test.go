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
		canRead        bool
		canWrite       bool
		expectedStatus int
	}{
		{true, true, true, http.StatusOK},
		{false, true, true, http.StatusForbidden},
	}

	for _, tt := range parametrizedTests {
		t.Run("", func(t *testing.T) {
			handlerToTest := configureDummyContextMiddleware(
				containers.NewStringSet(),
				containers.NewStringSet(),
				tt.isGlobalAdmin,
				tt.canRead,
				tt.canWrite,
				"",
			)(
				HasGlobalAdminPermission()(
					nextHandler(),
				),
			)
			req := httptest.NewRequest("GET", "http://testing", nil)
			responseRecorder := httptest.NewRecorder()
			handlerToTest.ServeHTTP(responseRecorder, req)
			if responseRecorder.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, responseRecorder.Code)
			}
		})
	}
}
