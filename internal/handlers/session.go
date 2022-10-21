package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nrc-no/notcore/internal/utils"
)

func HandleSession() http.Handler {

	type SessionResponse struct {
		Authenticated bool  `json:"authenticated"`
		ExpiresInMs   int64 `json:"expiresInMs"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		session, ok := utils.GetSession(ctx)
		if !ok {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if !session.IsAuthenticated() {
			ret := SessionResponse{
				Authenticated: false,
			}
			json.NewEncoder(w).Encode(ret)
			return
		} else {
			expiration := session.GetExpiration()
			expiresInMs := expiration.Sub(time.Now()).Milliseconds()
			ret := SessionResponse{
				Authenticated: true,
				ExpiresInMs:   expiresInMs,
			}
			json.NewEncoder(w).Encode(ret)
		}

	})
}
