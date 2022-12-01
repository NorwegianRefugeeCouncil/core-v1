package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

type TokenClaims struct {
	Sub    string   `json:"sub"`
	Iss    string   `json:"iss"`
	Email  string   `json:"email"`
	Groups []string `json:"groups"`
	Iat    int64    `json:"iat"`
	Exp    int64    `json:"exp"`
}

const AuthHeaderFormatJWT = "jwt"
const AuthHeaderFormatBearerToken = "bearer-token"

type IDTokenVerifier interface {
	Verify(ctx context.Context, rawIDToken string) (IDToken, error)
}

type idTokenVerifier struct {
	verifier *oidc.IDTokenVerifier
}

func NewIDTokenVerifier(verifier *oidc.IDTokenVerifier) IDTokenVerifier {
	return idTokenVerifier{
		verifier: verifier,
	}
}

func (i idTokenVerifier) Verify(ctx context.Context, rawIDToken string) (IDToken, error) {
	return i.verifier.Verify(ctx, rawIDToken)
}

type IDToken interface {
	Claims(v interface{}) error
}

func Authentication(
	authHeaderName,
	authHeaderFormat string,
	idTokenVerifier IDTokenVerifier,
	loginURL string,
) func(handler http.Handler) http.Handler {

	redirectToLogin := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			authPayload := r.Header.Get(authHeaderName)
			if len(authPayload) == 0 {
				l.Warn("missing authentication header")
				redirectToLogin(w, r)
				return
			}

			var rawIdToken string

			if authHeaderFormat == AuthHeaderFormatJWT {
				// This is useful for scenarios where the token is a plain JWT.
				// The auth header is in the format "<Header>: <token>"
				rawIdToken = r.Header.Get(authHeaderName)
			} else if authHeaderFormat == AuthHeaderFormatBearerToken {
				// This is useful for scenarios where the token is a bearer token.
				// And the auth header is in the format "<Header>: Bearer <jwt token>".
				bearerTokenParts := strings.Split(r.Header.Get(authHeaderName), " ")
				if len(bearerTokenParts) != 2 {
					l.Warn("invalid bearer token format. parts != 2")
					redirectToLogin(w, r)
					return
				}
				if len(bearerTokenParts[0]) != 6 && bearerTokenParts[0] != "Bearer" {
					l.Warn("invalid bearer token format. Does not start with 'Bearer'")
					redirectToLogin(w, r)
					return
				}
				rawIdToken = bearerTokenParts[1]
			} else {
				l.Warn("invalid auth header format", zap.String("authHeaderFormat", authHeaderFormat))
				redirectToLogin(w, r)
				return
			}

			idToken, err := idTokenVerifier.Verify(ctx, rawIdToken)
			if err != nil {
				l.Warn("failed to verify token", zap.Error(err))
				redirectToLogin(w, r)
				return
			}

			var tokenParts = strings.Split(rawIdToken, ".")
			var claimsPart = tokenParts[1]

			var tokenClaims TokenClaims
			if err := idToken.Claims(&tokenClaims); err != nil {
				l.Warn("failed to extract claims from token",
					zap.Error(err),
					zap.String("token", claimsPart))
				redirectToLogin(w, r)
				return
			}

			if err := validateTokenClaims(tokenClaims); err != nil {
				l.Warn("failed to validate token claims",
					zap.Error(err),
					zap.String("token", claimsPart),
				)
				redirectToLogin(w, r)
				return
			}

			session := auth.NewAuthenticatedSession(
				tokenClaims.Groups,
				tokenClaims.Email,
				tokenClaims.Iss,
				tokenClaims.Sub,
				time.Unix(tokenClaims.Exp, 0),
				time.Unix(tokenClaims.Iat, 0),
			)

			ctx = utils.WithSession(ctx, session)
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
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
	if claims.Groups == nil {
		return fmt.Errorf("token is missing groups")
	}
	if claims.Exp == 0 {
		return fmt.Errorf("token is missing expiration")
	}
	if claims.Iat == 0 {
		return fmt.Errorf("token is missing issued at")
	}
	return nil
}
