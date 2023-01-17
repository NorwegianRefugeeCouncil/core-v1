package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/nrc-no/notcore/internal/utils"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/server"
	"github.com/nrc-no/notcore/internal/server/middleware"
	"github.com/spf13/cobra"
)

const (
	envDbDSN                   = "CORE_DB_DSN"
	envDbDriver                = "CORE_DB_DRIVER"
	envListenAddress           = "CORE_LISTEN_ADDRESS"
	envLogoutURL               = "CORE_LOGOUT_URL"
	envLoginURL                = "CORE_LOGIN_URL"
	envTokenRefreshURL         = "CORE_TOKEN_REFRESH_URL"
	envTokenRefreshInterval    = "CORE_TOKEN_REFRESH_INTERVAL"
	envJwtGlobalAdminGroup     = "CORE_JWT_GLOBAL_ADMIN_GROUP"
	envJwtCanReadGroup         = "CORE_JWT_CAN_READ_GROUP"
	envJwtCanWriteGroup        = "CORE_JWT_CAN_WRITE_GROUP"
	envIdTokenHeaderName       = "CORE_ID_TOKEN_HEADER_NAME"
	envIdTokenHeaderFormat     = "CORE_ID_TOKEN_HEADER_FORMAT"
	envAccessTokenHeaderName   = "CORE_ACCESS_TOKEN_HEADER_NAME"
	envAccessTokenHeaderFormat = "CORE_ACCESS_TOKEN_HEADER_FORMAT"
	envOidcIssuerURL           = "CORE_OIDC_ISSUER"
	envOidcClientID            = "CORE_OAUTH_CLIENT_ID"
	envHashKey1                = "CORE_HASH_KEY_1"
	envBlockKey1               = "CORE_BLOCK_KEY_1"
	envHashKey2                = "CORE_HASH_KEY_2"
	envBlockKey2               = "CORE_BLOCK_KEY_2"

	flagDbDSN                   = "db-dsn"
	flagDbDriver                = "db-driver"
	flagListenAddress           = "listen-address"
	flagLogoutURL               = "logout-url"
	flagLoginURL                = "login-url"
	flagTokenRefreshURL         = "token-refresh-url"
	flagTokenRefreshInterval    = "token-refresh-interval"
	flagJwtGlobalAdminGroup     = "jwt-global-admin-group"
	flagJwtCanReadGroup         = "jwt-can-read-group"
	flagJwtCanWriteGroup        = "jwt-can-write-group"
	flagIdTokenHeaderName       = "id-token-header-name"
	flagIdTokenHeaderFormat     = "id-token-header-format"
	flagAccessTokenHeaderName   = "access-token-header-name"
	flagAccessTokenHeaderFormat = "access-token-header-format"
	flagOidcIssuerURL           = "oidc-issuer"
	flagOidcClientID            = "oauth-client-id"
	flagHashKey1                = "hash-key-1"
	flagBlockKey1               = "block-key-1"
	flagHashKey2                = "hash-key-2"
	flagBlockKey2               = "block-key-2"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)

		go func() {
			<-sig
			cancel()
		}()

		dbDsn := getFlagOrEnv(cmd, flagDbDSN, envDbDSN)
		if len(dbDsn) == 0 {
			return fmt.Errorf("--%s is required", flagDbDSN)
		}

		dbDriver := getFlagOrEnv(cmd, flagDbDriver, envDbDriver)
		if len(dbDriver) == 0 {
			return fmt.Errorf("--%s is required", flagDbDriver)
		}

		jwtGroupGlobalAdmin := getFlagOrEnv(cmd, flagJwtGlobalAdminGroup, envJwtGlobalAdminGroup)
		if len(jwtGroupGlobalAdmin) == 0 {
			return fmt.Errorf("--%s is required", flagJwtGlobalAdminGroup)
		}

		jwtGroupCanRead := getFlagOrEnv(cmd, flagJwtCanReadGroup, envJwtCanReadGroup)
		if len(jwtGroupCanRead) == 0 {
			return fmt.Errorf("--%s is required", flagJwtCanReadGroup)
		}

		jwtGroupCanWrite := getFlagOrEnv(cmd, flagJwtCanWriteGroup, envJwtCanWriteGroup)
		if len(jwtGroupCanWrite) == 0 {
			return fmt.Errorf("--%s is required", flagJwtCanWriteGroup)
		}

		listenAddress := getFlagOrEnv(cmd, flagListenAddress, envListenAddress)
		if len(listenAddress) == 0 {
			return fmt.Errorf("--%s is required", flagListenAddress)
		}

		logoutURL := getFlagOrEnv(cmd, flagLogoutURL, envLogoutURL)
		if len(logoutURL) == 0 {
			return fmt.Errorf("--%s is required", flagLogoutURL)
		}

		loginURL := getFlagOrEnv(cmd, flagLoginURL, envLoginURL)
		if len(logoutURL) == 0 {
			return fmt.Errorf("--%s is required", flagLoginURL)
		}

		refreshURL := getFlagOrEnv(cmd, flagTokenRefreshURL, envTokenRefreshURL)
		if len(refreshURL) == 0 {
			return fmt.Errorf("--%s is required", flagTokenRefreshURL)
		}

		tokenRefreshIntervalStr := cmd.Flag(flagTokenRefreshInterval).Value.String()
		if tokenRefreshIntervalStr == "0s" {
			tokenRefreshIntervalStr = os.Getenv(envTokenRefreshInterval)
			if tokenRefreshIntervalStr == "" {
				return fmt.Errorf("--%s is required", flagTokenRefreshInterval)
			}
		}
		tokenRefreshInterval, err := time.ParseDuration(tokenRefreshIntervalStr)
		if err != nil {
			return fmt.Errorf("--%s is invalid: %w", flagTokenRefreshInterval, err)
		}

		idTokenHeaderName := getFlagOrEnv(cmd, flagIdTokenHeaderName, envIdTokenHeaderName)
		if len(idTokenHeaderName) == 0 {
			return fmt.Errorf("--%s is required", flagIdTokenHeaderName)
		}

		idTokenHeaderFormat := getFlagOrEnv(cmd, flagIdTokenHeaderFormat, envIdTokenHeaderFormat)
		if len(idTokenHeaderFormat) == 0 {
			return fmt.Errorf("--%s is required", flagIdTokenHeaderFormat)
		}

		accessTokenHeaderName := getFlagOrEnv(cmd, flagAccessTokenHeaderName, envAccessTokenHeaderName)
		if len(idTokenHeaderName) == 0 {
			return fmt.Errorf("--%s is required", flagAccessTokenHeaderName)
		}

		accessTokenHeaderFormat := getFlagOrEnv(cmd, flagAccessTokenHeaderFormat, envAccessTokenHeaderFormat)
		if len(idTokenHeaderFormat) == 0 {
			return fmt.Errorf("--%s is required", flagAccessTokenHeaderFormat)
		}

		oidcIssuerURL := getFlagOrEnv(cmd, flagOidcIssuerURL, envOidcIssuerURL)
		if len(oidcIssuerURL) == 0 {
			return fmt.Errorf("--%s is required", flagOidcIssuerURL)
		}

		oauthClientID := getFlagOrEnv(cmd, flagOidcClientID, envOidcClientID)
		if len(oauthClientID) == 0 {
			return fmt.Errorf("--%s is required", flagOidcClientID)
		}

		hashKey1 := getFlagOrEnv(cmd, flagHashKey1, envHashKey1)
		if len(hashKey1) == 0 {
			return fmt.Errorf("--%s is required", flagHashKey1)
		}
		blockKey1 := getFlagOrEnv(cmd, flagBlockKey1, envBlockKey1)
		if len(blockKey1) == 0 {
			return fmt.Errorf("--%s is required", flagBlockKey1)
		}

		hashKey2 := getFlagOrEnv(cmd, flagHashKey2, envHashKey2)
		if len(hashKey2) == 0 {
			return fmt.Errorf("--%s is required", flagHashKey2)
		}
		blockKey2 := getFlagOrEnv(cmd, flagBlockKey2, envBlockKey2)
		if len(blockKey2) == 0 {
			return fmt.Errorf("--%s is required", flagBlockKey2)
		}

		options := server.Options{
			Address:              listenAddress,
			DatabaseDriver:       dbDriver,
			DatabaseDSN:          dbDsn,
			LogoutURL:            logoutURL,
			LoginURL:             loginURL,
			TokenRefreshURL:      refreshURL,
			TokenRefreshInterval: tokenRefreshInterval,
			JwtGroups: utils.JwtGroupOptions{
				GlobalAdmin: jwtGroupGlobalAdmin,
				CanRead:     jwtGroupCanRead,
				CanWrite:    jwtGroupCanWrite,
			},
			IdTokenAuthHeaderName:   idTokenHeaderName,
			IdTokenAuthHeaderFormat: idTokenHeaderFormat,
			AccessTokenHeaderName:   accessTokenHeaderName,
			AccessTokenHeaderFormat: accessTokenHeaderFormat,
			OIDCIssuerURL:           oidcIssuerURL,
			OAuthClientID:           oauthClientID,
			HashKey1:                hashKey1,
			BlockKey1:               blockKey1,
			HashKey2:                hashKey2,
			BlockKey2:               blockKey2,
		}

		srv, err := options.New(ctx)
		if err != nil {
			return err
		}

		if err := srv.Start(ctx); err != nil {
			if !errors.Is(err, net.ErrClosed) {
				return err
			}
			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String(flagListenAddress, "", cleanDoc(fmt.Sprintf(`
listen address. Can also be set with %[2]s
For example:
	--%[1]s=":8080"
	--%[1]s="127.0.0.1:8080"
`, flagListenAddress, envListenAddress)))

	serveCmd.PersistentFlags().String(flagDbDriver, "", cleanDoc(fmt.Sprintf(`
database driver. Can also be set with %s

Allowed values are 
	- sqlite (experimental support)
	- postgres
`, envDbDriver)))

	serveCmd.PersistentFlags().String(flagDbDSN, "", fmt.Sprintf("database dsn. Can also be set with %s", envDbDSN))

	serveCmd.PersistentFlags().String(flagLogoutURL, "", cleanDoc(fmt.Sprintf(`
logout url. Can also be set with %s

The URL is used to populate the "href" attribute of the logout button.
`, envLogoutURL)))

	serveCmd.PersistentFlags().String(flagLoginURL, "", cleanDoc(fmt.Sprintf(`
login url. Can also be set with %s

This URL is used to redirect the user with a malformed authentication header or token to the login page.
This is useful when the user token is expired, or if there is a problem with the authentication.
`, envLoginURL)))

	serveCmd.PersistentFlags().String(flagTokenRefreshURL, "", cleanDoc(fmt.Sprintf(`
session refresh url. Can also be set with %s

This URL is used to refresh the user session. It is called by the frontend application to refresh the user session
when it is about to expire.
`, envTokenRefreshURL)))

	serveCmd.PersistentFlags().String(flagJwtGlobalAdminGroup, "", cleanDoc(fmt.Sprintf(`
jwt global admin group. Can also be set with %s

This group is used to identify the global admin group. If the user is part of this group, he will be considered 
as having Global Administrator permissions. 
`, envJwtGlobalAdminGroup)))

	serveCmd.PersistentFlags().String(flagJwtCanReadGroup, "", cleanDoc(fmt.Sprintf(`
jwt can read group. Can also be set with %s

This group is used to identify the can read group. If the user is part of this group, he will be considered 
as having Read permissions. 
`, envJwtGlobalAdminGroup)))

	serveCmd.PersistentFlags().String(flagJwtCanWriteGroup, "", cleanDoc(fmt.Sprintf(`
jwt can write group. Can also be set with %s

This group is used to identify the can write group. If the user is part of this group, he will be considered 
as having Write permissions. 
`, envJwtGlobalAdminGroup)))

	serveCmd.PersistentFlags().String(flagIdTokenHeaderName, "", cleanDoc(fmt.Sprintf(`
id token header name. Can also be set with %[2]s

Different deployment scenarios exist for setting up an authentication proxy.
Usually, a proxy will forward the authorization information via a header. 
This parameter is used to specify the name of the header that contains the authorization information.

For example
	--%[1]s="X-Forwarded-ID-Token"
`, flagIdTokenHeaderName, envIdTokenHeaderName)))

	serveCmd.PersistentFlags().String(flagIdTokenHeaderFormat, "", cleanDoc(fmt.Sprintf(`
id token header format. Can also be set with %s. 
Allowed values are 
- "%s"
- "%s"

Depending on the authentication proxy, the value of the authentication header can be in different formats.
This parameter is used to specify the expected format of the authentication header.

A value of "%[2]s" means that the header value is a JWT token.
The application will expect a header like "<HeaderName>: <Token>".

A value of "%[3]s" means that the header value is a JWT token prefixed with "bearer".
The application will expect a header like "<HeaderName>: bearer <Token>".
`,
		envIdTokenHeaderFormat,
		middleware.AuthHeaderFormatJWT,
		middleware.AuthHeaderFormatBearerToken)))

	serveCmd.PersistentFlags().String(flagAccessTokenHeaderName, "", cleanDoc(fmt.Sprintf(`
access token header name. Can also be set with %[2]s

Different deployment scenarios exist for setting up an authentication proxy.
Usually, a proxy will forward the authorization information via a header. 
This parameter is used to specify the name of the header that contains the authorization information.

For example
	--%[1]s="X-Forwarded-Access-Token"
`, flagIdTokenHeaderName, envIdTokenHeaderName)))

	serveCmd.PersistentFlags().String(flagAccessTokenHeaderFormat, "", cleanDoc(fmt.Sprintf(`
access token header format. Can also be set with %s. 
Allowed values are 
- "%s"
- "%s"

Depending on the authentication proxy, the value of the authentication header can be in different formats.
This parameter is used to specify the expected format of the authentication header.

A value of "%[2]s" means that the header value is a JWT token.
The application will expect a header like "<HeaderName>: <Token>".

A value of "%[3]s" means that the header value is a JWT token prefixed with "bearer".
The application will expect a header like "<HeaderName>: bearer <Token>".
`,
		envIdTokenHeaderFormat,
		middleware.AuthHeaderFormatJWT,
		middleware.AuthHeaderFormatBearerToken)))

	serveCmd.PersistentFlags().String(flagOidcIssuerURL, "", cleanDoc(fmt.Sprintf(`
oidc issuer URL. Can also be set with %s

The oidc issuer URL is used to identify the OIDC provider. It is also used to retrieve the 
OIDC provider's discovery document'
`, envOidcIssuerURL)))

	serveCmd.PersistentFlags().String(flagOidcClientID, "", fmt.Sprintf("oauth client id. Can also be set with %s", envOidcClientID))

	serveCmd.PersistentFlags().Duration(flagTokenRefreshInterval, 0, cleanDoc(fmt.Sprintf(`
This flag specifies the interval at which user token should be refreshed. Can also be set with %s

For example, if the value of this flag is set to 50m, the token will be refreshed every 50 minutes.
The browser will be responsible for refreshing the token. So if the user does not have a browser window
opened, the token will not be refreshed.
`, envTokenRefreshInterval)))

	serveCmd.PersistentFlags().String(flagHashKey1, "", cleanDoc(fmt.Sprintf(`
This flag specifies the hex-encoded first hash key used to encrypt the session cookie. Can also be set with %s
`, envHashKey1)))

	serveCmd.PersistentFlags().String(flagBlockKey1, "", cleanDoc(fmt.Sprintf(`
This flag specifies the hex-encoded first block key used to encrypt the session cookie. Can also be set with %s
`, envBlockKey1)))

	serveCmd.PersistentFlags().String(flagHashKey2, "", cleanDoc(fmt.Sprintf(`
This flag specifies the hex-encoded second hash key used to encrypt the session cookie. 
The second hash key is used to perform smooth key rotation.
Usually, the first hash key is moved to the second hash key, and a new first hash key is generated.
Can also be set with %s
`, envHashKey2)))

	serveCmd.PersistentFlags().String(flagBlockKey2, "", cleanDoc(fmt.Sprintf(`
This flag specifies the hex-encoded second block key used to encrypt the session cookie. 
The second block key is used to perform smooth key rotation.
Usually, the first block key is moved to the second block key, and a new first block key is generated.
Can also be set with %s
`, envBlockKey2)))

}

func cleanDoc(s string) string {
	if strings.HasPrefix(s, "\n") {
		s = s[1:]
	}
	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}
	return s
}

func getFlagOrEnv(cmd *cobra.Command, flagName string, envName string) string {
	flagValue := cmd.Flag(flagName).Value.String()
	if len(flagValue) > 0 {
		return flagValue
	}
	return os.Getenv(envName)
}
