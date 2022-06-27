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

		dsn := cmd.Flag("db-dsn").Value.String()
		fileDns := cmd.Flag("db-dsn-file").Value.String()
		if len(dsn) == 0 {
			if len(fileDns) == 0 {
				fmt.Println("db-dsn or db-dsn-file is required")
				return cmd.Usage()
			}
			dsnBytes, err := os.ReadFile(fileDns)
			if err != nil {
				return err
			}
			dsn = string(dsnBytes)
		}

		options := server.Options{
			Address:        cmd.Flag("listen-address").Value.String(),
			DatabaseDriver: cmd.Flag("db-driver").Value.String(),
			DatabaseDSN:    dsn,
		}

		srv, err := options.New()
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
	serveCmd.PersistentFlags().String("listen-address", ":8080", "listen address")
	serveCmd.PersistentFlags().String("db-driver", "", "database driver")
	serveCmd.PersistentFlags().String("db-dsn", "", "database dsn")
	serveCmd.PersistentFlags().String("db-dsn-file", "", "database dsn file")
}
