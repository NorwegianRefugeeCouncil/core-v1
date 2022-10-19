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
	envIDTokenHeaderName   = "CORE_ID_TOKEN_HEADER_NAME"

	flagDbDSN               = "db-dsn"
	flagDbDriver            = "db-driver"
	flagListenAddress       = "listen-address"
	flagLogoutURL           = "logout-url"
	flagJwtGlobalAdminGroup = "jwt-global-admin-group"
	flagIDTokenHeaderName   = "id-token-header-name"
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

		idTokenHeaderName := getFlagOrEnv(cmd, flagIDTokenHeaderName, envIDTokenHeaderName)
		if len(idTokenHeaderName) == 0 {
			return fmt.Errorf("--%s is required", flagIDTokenHeaderName)
		}

		options := server.Options{
			Address:             listenAddress,
			DatabaseDriver:      dbDriver,
			DatabaseDSN:         dbDsn,
			LogoutURL:           logoutURL,
			JwtGroupGlobalAdmin: jwtGroupGlobalAdmin,
			IDTokenHeaderName:   idTokenHeaderName,
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
	serveCmd.PersistentFlags().String(flagIDTokenHeaderName, "", fmt.Sprintf("id token header name. Can also be set with %s", envIDTokenHeaderName))
}

func getFlagOrEnv(cmd *cobra.Command, flagName string, envName string) string {
	flagValue := cmd.Flag(flagName).Value.String()
	if len(flagValue) > 0 {
		return flagValue
	}
	return os.Getenv(envName)
}
