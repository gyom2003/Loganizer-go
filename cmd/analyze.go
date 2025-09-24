package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
)

var (
	configPath string
	outputPath string
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze log files based on configuration",
	Long: `Analyze multiple log files concurrently based on a JSON configuration file.

The analyze command reads a JSON configuration file containing log file definitions
and processes them in parallel using goroutines. Results are collected and can be
exported to a JSON report file.

Example:
  loganalyzer analyze --config config.json
  loganalyzer analyze -c config.json -o report.json`,
	RunE: runAnalyze,
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	// Validate required flags
	if configPath == "" {
		return fmt.Errorf("config file path is required (use --config or -c flag)")
	}

	// Load configuration
	fmt.Printf("Loading configuration from: %s\n", configPath)
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		// Handle custom parse error
		var parseErr *config.ParseError
		if errors.As(err, &parseErr) {
			return fmt.Errorf("configuration parsing failed: %w", parseErr)
		}
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	fmt.Printf("Found %d log(s) to analyze\n", len(cfg))

	// Create analyzer and process logs
	fmt.Println("Starting concurrent log analysis...")
	analyzer := analyzer.NewAnalyzer()
	results := analyzer.AnalyzeAll(cfg)

	// Create reporter and display results
	reporter := reporter.NewReporter()
	reporter.PrintSummary(results)

	// Export to JSON if output path is specified
	if outputPath != "" {
		fmt.Printf("\nExporting results to: %s\n", outputPath)
		if err := reporter.ExportToJSON(results, outputPath); err != nil {
			return fmt.Errorf("failed to export results: %w", err)
		}
		fmt.Println("Results exported successfully!")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Define flags
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to JSON configuration file (required)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to output JSON report file (optional)")

	// Mark config flag as required
	analyzeCmd.MarkFlagRequired("config")
}
