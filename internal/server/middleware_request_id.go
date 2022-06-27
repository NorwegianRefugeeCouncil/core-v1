package server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nrc-no/notcore/internal/utils"
)

func requestIdMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}
		ctx := r.Context()
		ctx = utils.WithRequestID(ctx, requestID)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
