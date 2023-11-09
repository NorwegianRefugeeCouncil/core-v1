package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type TokenClaims struct {
	Sub   string `json:"sub"`
	Iss   string `json:"iss"`
	Email string `json:"email"`
	Iat   int64  `json:"iat"`
	Exp   int64  `json:"exp"`
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
	idTokenHeaderName,
	idTokenHeaderFormat,
	accessTokenHeaderName,
	accessTokenHeaderFormat string,
	provider *oidc.Provider,
	idTokenVerifier IDTokenVerifier,
	sessionStore *sessions.CookieStore,
	loginURL string,
) func(handler http.Handler) http.Handler {

	redirectToLogin := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			l := logging.NewLogger(ctx)

			authPayload := r.Header.Get(idTokenHeaderName)
			if len(authPayload) == 0 {
				l.Warn("missing authentication header")
				redirectToLogin(w, r)
				return
			}

			var rawIdToken string

			rawIdToken, err := parseAuthHeader(r, idTokenHeaderName, idTokenHeaderFormat)
			if err != nil {
				l.Warn("invalid authentication header", zap.Error(err))
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

			userInfoClaims, err := getUserInfoClaims(sessionStore, provider, r, w, accessTokenHeaderName, accessTokenHeaderFormat)
			if err != nil {
				l.Warn("failed to get user info claims", zap.Error(err))
				redirectToLogin(w, r)
				return
			}

			userSession := auth.NewAuthenticatedSession(
				userInfoClaims.Groups,
				tokenClaims.Email,
				tokenClaims.Iss,
				tokenClaims.Sub,
				time.Unix(tokenClaims.Exp, 0),
				time.Unix(tokenClaims.Iat, 0),
			)

			ctx = utils.WithSession(ctx, userSession)
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}

type userInfoClaims struct {
	Groups          []string `json:"groups"`
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
	if claims.Exp == 0 {
		return fmt.Errorf("token is missing expiration")
	}
	if claims.Iat == 0 {
		return fmt.Errorf("token is missing issued at")
	}
	return nil
}

func getUserInfoClaims(
	sessionStore sessions.Store,
	provider *oidc.Provider,
	r *http.Request,
	w http.ResponseWriter,
	accessTokenHeaderName,
	accessTokenHeaderFormat string,
) (*userInfoClaims, error) {

	ctx := r.Context()

	cookieSession, _ := sessionStore.Get(r, "core-session")

	userInfoInterface, ok := cookieSession.Values["oidc_userinfo"]
	if ok {
		userInfoBytesRead, ok := userInfoInterface.([]byte)
		userInfo := userInfoClaims{}
		json.Unmarshal(userInfoBytesRead, &userInfo)
		if ok {
			return &userInfo, nil
		}
	}

	rawAccessToken, err := parseAuthHeader(r, accessTokenHeaderName, accessTokenHeaderFormat)
	if err != nil {
		return nil, err
	}

	userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: rawAccessToken}))
	if err != nil {
		return nil, err
	}

	var info userInfoClaims
	if err := userInfo.Claims(&info); err != nil {
		return nil, err
	}

	if info.Groups == nil {
		return nil, fmt.Errorf("missing groups claim")
	}

	infoBytesWrite, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	cookieSession.Values["oidc_userinfo"] = infoBytesWrite
	if err := cookieSession.Save(r, w); err != nil {
		return nil, err
	}

	return &info, nil
}

func parseAuthHeader(r *http.Request, headerName, headerFormat string) (string, error) {
	var rawToken string
	if headerFormat == AuthHeaderFormatJWT {
		// This is useful for scenarios where the token is a plain JWT.
		// The auth header is in the format "<Header>: <token>"
		rawToken = r.Header.Get(headerName)
	} else if headerFormat == AuthHeaderFormatBearerToken {
		// This is useful for scenarios where the token is a bearer token.
		// And the auth header is in the format "<Header>: Bearer <jwt token>".
		bearerTokenParts := strings.Split(r.Header.Get(headerName), " ")
		if len(bearerTokenParts) != 2 {
			return "", errors.New("invalid bearer token format. parts != 2")
		}
		if len(bearerTokenParts[0]) != 6 && bearerTokenParts[0] != "Bearer" {
			return "", errors.New("invalid bearer token format. Does not start with 'Bearer'")
		}
		rawToken = bearerTokenParts[1]
	} else {
		return "", fmt.Errorf("invalid auth header format: %s", headerFormat)
	}
	return rawToken, nil
}
