package cmd

import (
	"fmt"
	"os"

	"github.com/nrc-no/notcore/internal/logging"
	"github.com/spf13/cobra"
)

const (
	flagLogLevel = "log-level"
	envLogLevel  = "CORE_LOG_LEVEL"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "core",
	Short: "",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		logLevel := getFlagOrEnv(cmd, flagLogLevel, envLogLevel)
		if logLevel == "" {
			logLevel = string(logging.DefaultLevel)
		}
		if err := logging.SetLogLevel(logLevel); err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String(flagLogLevel, "", cleanDoc(fmt.Sprintf(`
Log Level. Valid values are: %v. Can also be set with the %s environment variable.
`,
		logging.AllowedLogLevels(),
		envLogLevel)))
}
