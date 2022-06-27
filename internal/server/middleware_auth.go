package server

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/utils"
)

func authMiddleware(userRepo db.UserRepo) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		type authHeaderClaims struct {
			Sub    string `json:"sub"`
			Email  string `json:"email"`
			Issuer string `json:"iss"`
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeaderBase64 := r.Header.Get("X-Jwt-Payload")
			if len(authHeaderBase64) == 0 {
				http.Error(w, "Invalid authorization header", http.StatusBadRequest)
				return
			}

			authHeaderJsonBytes, err := base64.RawURLEncoding.DecodeString(authHeaderBase64)
			if err != nil {
				http.Error(w, "Invalid authorization header: "+err.Error(), http.StatusBadRequest)
				return
			}

			var payload authHeaderClaims
			if err := json.Unmarshal(authHeaderJsonBytes, &payload); err != nil {
				http.Error(w, "Invalid authorization header: "+err.Error(), http.StatusBadRequest)
				return
			}

			user, err := userRepo.Put(r.Context(), &api.User{
				ID:      "",
				Subject: payload.Sub,
				Email:   payload.Email,
			})
			if err != nil {
				http.Error(w, "couldn't save user: "+err.Error(), http.StatusInternalServerError)
				return
			}

			ctx := r.Context()
			ctx = utils.WithUser(ctx, *user)
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}
