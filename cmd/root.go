// Package cmd contains the CLI commands for loganalyzer
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "A distributed log analysis tool",
	Long: `LogAnalyzer is a CLI tool designed to help system administrators 
analyze log files from various sources (servers, applications) in parallel.

The tool provides centralized analysis of multiple logs simultaneously 
while robustly handling potential errors.

Example usage:
  loganalyzer analyze --config config.json --output report.json`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add any global flags here if needed
}
