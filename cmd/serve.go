package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/nrc-no/notcore/internal/server"
	"github.com/spf13/cobra"
)

const (
	envDbDSN               = "CORE_DB_DSN"
	envDbDriver            = "CORE_DB_DRIVER"
	envListenAddress       = "CORE_LISTEN_ADDRESS"
	envLogoutURL           = "CORE_LOGOUT_URL"
	envRefreshTokenURL     = "CORE_REFRESH_TOKEN_URL"
	envRefreshTokenBefore  = "CORE_REFRESH_TOKEN_BEFORE"
	envJwtGlobalAdminGroup = "CORE_JWT_GLOBAL_ADMIN_GROUP"
	envAuthHeaderName      = "CORE_AUTH_HEADER_NAME"
	envAuthHeaderFormat    = "CORE_AUTH_HEADER_FORMAT"
	envOidcIssuer          = "CORE_OIDC_ISSUER"
	envOidcClientID        = "CORE_OAUTH_CLIENT_ID"

	flagDbDSN               = "db-dsn"
	flagDbDriver            = "db-driver"
	flagListenAddress       = "listen-address"
	flagLogoutURL           = "logout-url"
	flagRefreshTokenURL     = "refresh-token-url"
	flagRefreshTokenBefore  = "refresh-token-before"
	flagJwtGlobalAdminGroup = "jwt-global-admin-group"
	flagAuthHeaderName      = "auth-header-name"
	flagAuthHeaderFormat    = "auth-header-format"
	flagOidcIssuer          = "oidc-issuer"
	flagOidcClientID        = "oauth-client-id"
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

		listenAddress := getFlagOrEnv(cmd, flagListenAddress, envListenAddress)
		if len(listenAddress) == 0 {
			return fmt.Errorf("--%s is required", flagListenAddress)
		}

		logoutURL := getFlagOrEnv(cmd, flagLogoutURL, envLogoutURL)
		if len(logoutURL) == 0 {
			return fmt.Errorf("--%s is required", flagLogoutURL)
		}

		refreshURL := getFlagOrEnv(cmd, flagRefreshTokenURL, envRefreshTokenURL)
		if len(refreshURL) == 0 {
			return fmt.Errorf("--%s is required", flagRefreshTokenURL)
		}

		refreshTokenBefore := getFlagOrEnv(cmd, flagRefreshTokenBefore, envRefreshTokenBefore)
		if refreshTokenBefore == "0s" {
			return fmt.Errorf("--%s is required", flagRefreshTokenBefore)
		}
		refreshBeforeDuration, err := time.ParseDuration(refreshTokenBefore)
		if err != nil {
			return fmt.Errorf("--%s is invalid: %w", flagRefreshTokenBefore, err)
		}

		authHeaderName := getFlagOrEnv(cmd, flagAuthHeaderName, envAuthHeaderName)
		if len(authHeaderName) == 0 {
			return fmt.Errorf("--%s is required", flagAuthHeaderName)
		}

		authHeaderFormat := getFlagOrEnv(cmd, flagAuthHeaderFormat, envAuthHeaderFormat)
		if len(authHeaderFormat) == 0 {
			return fmt.Errorf("--%s is required", flagAuthHeaderFormat)
		}

		oidcIssuer := getFlagOrEnv(cmd, flagOidcIssuer, envOidcIssuer)
		if len(oidcIssuer) == 0 {
			return fmt.Errorf("--%s is required", flagOidcIssuer)
		}

		oauthClientID := getFlagOrEnv(cmd, flagOidcClientID, envOidcClientID)
		if len(oauthClientID) == 0 {
			return fmt.Errorf("--%s is required", flagOidcClientID)
		}

		options := server.Options{
			Address:             listenAddress,
			DatabaseDriver:      dbDriver,
			DatabaseDSN:         dbDsn,
			LogoutURL:           logoutURL,
			RefreshTokenURL:     refreshURL,
			RefreshTokenBefore:  refreshBeforeDuration,
			JwtGroupGlobalAdmin: jwtGroupGlobalAdmin,
			AuthHeaderName:      authHeaderName,
			AuthHeaderFormat:    authHeaderFormat,
			OIDCIssuerURL:       oidcIssuer,
			OAuthClientID:       oauthClientID,
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
	serveCmd.PersistentFlags().String(flagListenAddress, "", fmt.Sprintf("listen address. Can also be set with %s", envListenAddress))
	serveCmd.PersistentFlags().String(flagDbDriver, "", fmt.Sprintf("database driver. Can also be set with %s", envDbDriver))
	serveCmd.PersistentFlags().String(flagDbDSN, "", fmt.Sprintf("database dsn. Can also be set with %s", envDbDSN))
	serveCmd.PersistentFlags().String(flagLogoutURL, "", fmt.Sprintf("logout url. Can also be set with %s", envLogoutURL))
	serveCmd.PersistentFlags().String(flagRefreshTokenURL, "", fmt.Sprintf("session refresh url. Can also be set with %s", envRefreshTokenURL))
	serveCmd.PersistentFlags().String(flagJwtGlobalAdminGroup, "", fmt.Sprintf("jwt global admin group. Can also be set with %s", envJwtGlobalAdminGroup))
	serveCmd.PersistentFlags().String(flagAuthHeaderName, "", fmt.Sprintf("auth header name. Can also be set with %s", envAuthHeaderName))
	serveCmd.PersistentFlags().String(flagAuthHeaderFormat, "", fmt.Sprintf(`auth header format. Can also be set with %s. Allowed values are "%s", "%s"`,
		envAuthHeaderFormat,
		server.AuthHeaderFormatJWT,
		server.AuthHeaderFormatBearerToken))
	serveCmd.PersistentFlags().String(flagOidcIssuer, "", fmt.Sprintf("oidc issuer. Can also be set with %s", envOidcIssuer))
	serveCmd.PersistentFlags().String(flagOidcClientID, "", fmt.Sprintf("oauth client id. Can also be set with %s", envOidcClientID))
	serveCmd.PersistentFlags().Duration(flagRefreshTokenBefore, 0, fmt.Sprintf("refresh token before. Can also be set with %s", envRefreshTokenBefore))
}

func getFlagOrEnv(cmd *cobra.Command, flagName string, envName string) string {
	flagValue := cmd.Flag(flagName).Value.String()
	if len(flagValue) > 0 {
		return flagValue
	}
	return os.Getenv(envName)
}
