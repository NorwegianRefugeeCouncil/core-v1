package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/nrc-no/notcore/internal/server"
	"github.com/spf13/cobra"
)

const (
	envDbDSN               = "CORE_DB_DSN"
	envDbDriver            = "CORE_DB_DRIVER"
	envListenAddress       = "CORE_LISTEN_ADDRESS"
	envLogoutURL           = "CORE_LOGOUT_URL"
	envJwtGlobalAdminGroup = "CORE_JWT_GLOBAL_ADMIN_GROUP"
	envAuthHeaderName      = "CORE_AUTH_HEADER_NAME"
	envAuthHeaderFormat    = "CORE_AUTH_HEADER_FORMAT"

	flagDbDSN               = "db-dsn"
	flagDbDriver            = "db-driver"
	flagListenAddress       = "listen-address"
	flagLogoutURL           = "logout-url"
	flagJwtGlobalAdminGroup = "jwt-global-admin-group"
	flagAuthHeaderName      = "auth-header-name"
	flagAuthHeaderFormat    = "auth-header-format"
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

		authHeaderName := getFlagOrEnv(cmd, flagAuthHeaderName, envAuthHeaderName)
		if len(authHeaderName) == 0 {
			return fmt.Errorf("--%s is required", flagAuthHeaderName)
		}

		authHeaderFormat := getFlagOrEnv(cmd, flagAuthHeaderFormat, envAuthHeaderFormat)
		if len(authHeaderFormat) == 0 {
			return fmt.Errorf("--%s is required", flagAuthHeaderFormat)
		}

		options := server.Options{
			Address:             listenAddress,
			DatabaseDriver:      dbDriver,
			DatabaseDSN:         dbDsn,
			LogoutURL:           logoutURL,
			JwtGroupGlobalAdmin: jwtGroupGlobalAdmin,
			AuthHeaderName:      authHeaderName,
			AuthHeaderFormat:    authHeaderFormat,
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
	serveCmd.PersistentFlags().String(flagJwtGlobalAdminGroup, "", fmt.Sprintf("jwt global admin group. Can also be set with %s", envJwtGlobalAdminGroup))
	serveCmd.PersistentFlags().String(flagAuthHeaderName, "", fmt.Sprintf("auth header name. Can also be set with %s", envAuthHeaderName))
	serveCmd.PersistentFlags().String(flagAuthHeaderFormat, "", fmt.Sprintf(`auth header format. Can also be set with %s. Allowed values are "%s", "%s" and "%s"`,
		envAuthHeaderFormat,
		server.AuthHeaderFormatJWT,
		server.AuthHeaderFormatBearerToken,
		server.AuthHeaderFormatJsonBase64UrlEncodedClaims))
}

func getFlagOrEnv(cmd *cobra.Command, flagName string, envName string) string {
	flagValue := cmd.Flag(flagName).Value.String()
	if len(flagValue) > 0 {
		return flagValue
	}
	return os.Getenv(envName)
}
