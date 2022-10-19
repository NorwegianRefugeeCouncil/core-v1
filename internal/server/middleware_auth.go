package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
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

const AuthHeaderFormatJsonBase64UrlEncodedClaims = "json-base64-url-encoded-claims"
const AuthHeaderFormatJWT = "jwt"
const AuthHeaderFormatBearerToken = "bearer-token"

func authMiddleware(authHeaderName, authHeaderFormat string) func(handler http.Handler) http.Handler {

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			authPayload := r.Header.Get(authHeaderName)
			if len(authPayload) == 0 {
				http.Error(w, "Invalid authorization header", http.StatusBadRequest)
				return
			}

			var claims TokenClaims

			if authHeaderFormat == AuthHeaderFormatJsonBase64UrlEncodedClaims {
				// This is useful for envoy-based scenarios where the token claims
				// are already parsed and base64 encoded in the header.
				var err error
				claims, err = extractJsonBase64UrlEncodedTokenClaims(authPayload)
				if err != nil {
					l.Warn("failed to extract json-base64-urlencoded claims", zap.Error(err))
					http.Error(w, "Invalid authorization header", http.StatusBadRequest)
					return
				}
			} else if authHeaderFormat == AuthHeaderFormatJWT {
				// This is useful for scenarios where the token is a JWT.
				// Since we rely on the proxy to validate the token, we only
				// need to extract the claims, without validating the signature.
				// At some point we may want to validate the signature here,
				// as well as the authorized signature algorithms.
				// But we would need to get the public key from the issuer.
				var err error
				claims, err = extractJwtTokenClaims(authPayload)
				if err != nil {
					l.Warn("failed to extract jwt claims", zap.Error(err))
					http.Error(w, "Invalid authorization header", http.StatusBadRequest)
					return
				}
			} else if authHeaderFormat == AuthHeaderFormatBearerToken {
				// This is useful for scenarios where the token is a bearer token.
				// And the auth header is in the format "Bearer <jwt token>".
				var err error
				claims, err = extractBearerTokenClaims(authPayload)
				if err != nil {
					l.Warn("failed to extract jwt claims", zap.Error(err))
					http.Error(w, "Invalid authorization header", http.StatusBadRequest)
					return
				}
			} else {
				l.Error("invalid auth header format", zap.String("authHeaderFormat", authHeaderFormat))
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

// extractBearerTokenClaims will extract the claims from a bearer token.
// A bearer token header is in the format "Bearer <jwt token>".
func extractBearerTokenClaims(payload string) (TokenClaims, error) {
	bearerTokenParts := strings.Split(payload, " ")
	if len(bearerTokenParts) != 2 {
		return TokenClaims{}, fmt.Errorf("invalid bearer token format")
	}
	if bearerTokenParts[0] != "Bearer" {
		return TokenClaims{}, fmt.Errorf("invalid bearer token format")
	}
	var err error
	var claims TokenClaims
	claims, err = extractJwtTokenClaims(bearerTokenParts[1])
	if err != nil {
		return TokenClaims{}, err
	}
	return claims, nil
}

// extractJwtTokenClaims will extract the claims from a JWT token.
// The token is not validated, only the claims are extracted.
func extractJwtTokenClaims(payload string) (TokenClaims, error) {
	var jwtClaims jwt.MapClaims
	jwtParser := jwt.NewParser()
	_, _, err := jwtParser.ParseUnverified(payload, &jwtClaims)
	if err != nil {
		return TokenClaims{}, fmt.Errorf("failed to extract claims: %w", err)
	}
	if err := jwtClaims.Valid(); err != nil {
		return TokenClaims{}, fmt.Errorf("failed to extract claims: invalid token")
	}
	var ret TokenClaims
	if ret, err = unmarshalJwtClaims(jwtClaims); err != nil {
		return TokenClaims{}, fmt.Errorf("failed to extract claims: %w", err)
	}
	return ret, nil
}

// extractJsonBase64UrlEncodedTokenClaims will extract the claims from a base64 url encoded json string.
func extractJsonBase64UrlEncodedTokenClaims(payload string) (TokenClaims, error) {
	authHeaderJsonBytes, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return TokenClaims{}, fmt.Errorf("failed to extract claims: %w", err)
	}
	claims, err := unmarshalJsonClaims(authHeaderJsonBytes)
	if err != nil {
		return TokenClaims{}, fmt.Errorf("failed to extract claims: %w", err)
	}
	return claims, nil
}

// unmarshalJsonClaims will unmarshal the claims from a json string.
func unmarshalJsonClaims(jsonPayload []byte) (TokenClaims, error) {
	var claims TokenClaims
	if err := json.Unmarshal(jsonPayload, &claims); err != nil {
		return TokenClaims{}, err
	}
	return claims, nil
}

// unmarshalJwtClaims will unmarshal the claims from a jwt.MapClaims.
func unmarshalJwtClaims(jwtClaims jwt.MapClaims) (TokenClaims, error) {
	var claims TokenClaims
	for k, v := range jwtClaims {
		switch k {
		case "iss":
			if iss, ok := v.(string); ok {
				claims.Iss = iss
				break
			} else {
				return TokenClaims{}, fmt.Errorf("token is missing issuer")
			}
		case "sub":
			if sub, ok := v.(string); ok {
				claims.Sub = sub
				break
			} else {
				return TokenClaims{}, fmt.Errorf("token is missing subject")
			}
		case "email":
			if email, ok := v.(string); ok {
				claims.Email = email
				break
			} else {
				return TokenClaims{}, fmt.Errorf("token is missing email")
			}
		case "groups":
			groupsIntf, ok := v.([]interface{})
			if !ok {
				return TokenClaims{}, fmt.Errorf("token is missing groups")
			}
			for _, groupIntf := range groupsIntf {
				group, ok := groupIntf.(string)
				if !ok {
					return TokenClaims{}, fmt.Errorf("token is missing groups")
				}
				claims.Groups = append(claims.Groups, group)
			}
			break
		}
	}
	return claims, nil
}

// validateTokenClaims will validate the claims of a token.
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
