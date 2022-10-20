/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/nrc-no/notcore/cmd/mockdata"
	"github.com/spf13/cobra"
)

// mockDataCmd represents the mockData command
var mockDataCmd = &cobra.Command{
	Use:   "mock-data",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		countryID, err := cmd.Flags().GetString("country-id")
		if err != nil {
			return err
		}
		if countryID == "" {
			return fmt.Errorf("country-id is required")
		}

		count, err := cmd.Flags().GetUint("count")
		if err != nil {
			return err
		}
		if count == 0 {
			return fmt.Errorf("count is required")
		}

		return mockdata.Generate(countryID, count)
	},
}

func init() {
	rootCmd.AddCommand(mockDataCmd)
	mockDataCmd.PersistentFlags().String("country-id", "", "Country ID")
	mockDataCmd.PersistentFlags().Uint("count", 0, "Number of records to generate")
}
