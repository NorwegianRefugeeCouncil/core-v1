package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

type TokenClaims struct {
	Sub    string   `json:"sub"`
	Iss    string   `json:"iss"`
	Email  string   `json:"email"`
	Groups []string `json:"groups"`
}

func jwtMiddleware() func(handler http.Handler) http.Handler {

	const (
		keyJwtPayload = "X-Jwt-Payload"
	)

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			authHeaderBase64 := r.Header.Get(keyJwtPayload)
			if len(authHeaderBase64) == 0 {
				http.Error(w, "Invalid authorization header", http.StatusBadRequest)
				return
			}

			authHeaderJsonBytes, err := base64.RawURLEncoding.DecodeString(authHeaderBase64)
			if err != nil {
				http.Error(w, "Invalid authorization header: "+err.Error(), http.StatusBadRequest)
				return
			}

			claims, err := unmarshalTokenClaims(authHeaderJsonBytes)
			if err != nil {
				l.Error("failed to unmarshal token claims", zap.Error(err))
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			if err := validateTokenClaims(claims); err != nil {
				l.Error("failed to validate token claims", zap.Error(err))
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			user := &api.User{
				ID:      fmt.Sprintf("%s/%s", claims.Iss, claims.Sub),
				Issuer:  claims.Iss,
				Subject: claims.Sub,
				Email:   claims.Email,
				Groups:  claims.Groups,
			}

			ctx = utils.WithUser(ctx, *user)
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}

func unmarshalTokenClaims(payload []byte) (TokenClaims, error) {
	var claims TokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return TokenClaims{}, err
	}
	return claims, nil
}

func validateTokenClaims(claims TokenClaims) error {
	if claims.Iss == "" {
		return fmt.Errorf("token is missing issuer")
	}
	if claims.Sub == "" {
		return fmt.Errorf("token is missing subject")
	}
	if claims.Email == "" {
		return fmt.Errorf("token is missing email")
	}
	if len(claims.Groups) == 0 {
		return fmt.Errorf("token is missing groups")
	}
	return nil
}
