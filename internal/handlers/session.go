package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nrc-no/notcore/internal/utils"
)

func HandleSession() http.Handler {

	type SessionResponse struct {
		IsAuthenticated bool  `json:"isAuthenticated"`
		ExpiresInMs     int64 `json:"expiresInMs"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		session, ok := utils.GetSession(ctx)
		if !ok {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		ret := SessionResponse{
			IsAuthenticated: session.IsAuthenticated(),
		}

		if session.IsAuthenticated() {
			ret.ExpiresInMs = session.GetExpiration().Sub(time.Now()).Milliseconds()
		}

		json.NewEncoder(w).Encode(ret)

	})
}
