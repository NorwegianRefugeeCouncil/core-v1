package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "codegen",
	RunE: func(cmd *cobra.Command, args []string) error {

		inFile := os.Getenv("GOFILE")
		logger := log.New(os.Stdout, "[codegen] ", log.LstdFlags)

		for i, arg := range args {
			logger.Printf("arg %d: %s", i, arg)
		}
		structName, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		if structName == "" {
			return fmt.Errorf("struct name is required")
		}
		outputFileName, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		if outputFileName == "" {
			outputFileName = strings.ToLower(structName) + "_fields.go"
		}
		tagName, err := cmd.Flags().GetString("tag")
		if err != nil {
			return err
		}
		if tagName == "" {
			tagName = "json"
		}
		opts := Options{
			Logger:     logger,
			FileName:   inFile,
			StructName: structName,
			OutputFile: outputFileName,
			TagName:    tagName,
		}
		return gen(opts)
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
	rootCmd.Flags().String("type", "", "Struct name")
	rootCmd.Flags().String("output", "", "Output file")
	rootCmd.Flags().String("tag", "", "Tag name")
}
