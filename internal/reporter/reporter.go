// Package reporter handles exporting analysis results to JSON files
package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
)

// Reporter handles exporting results to JSON format
type Reporter struct{}

// NewReporter creates a new Reporter instance
func NewReporter() *Reporter {
	return &Reporter{}
}

// ExportToJSON exports analysis results to a JSON file
// If the output path includes directories that don't exist, they will be created
func (r *Reporter) ExportToJSON(results []analyzer.Result, outputPath string) error {
	// Create directories if they don't exist (bonus feature)
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Marshal results to JSON with pretty formatting
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal results to JSON: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

// PrintSummary prints a summary of results to the console
func (r *Reporter) PrintSummary(results []analyzer.Result) {
	fmt.Println("\n=== Log Analysis Summary ===")
	fmt.Printf("Total logs analyzed: %d\n\n", len(results))

	for _, result := range results {
		fmt.Printf("ID: %s\n", result.LogID)
		fmt.Printf("Path: %s\n", result.FilePath)
		fmt.Printf("Status: %s\n", result.Status)
		fmt.Printf("Message: %s\n", result.Message)
		if result.ErrorDetails != "" {
			fmt.Printf("Error: %s\n", result.ErrorDetails)
		}
		fmt.Println("---")
	}

	// Print statistics
	successful := 0
	failed := 0
	for _, result := range results {
		if result.Status == "OK" {
			successful++
		} else {
			failed++
		}
	}

	fmt.Printf("\nResults: %d successful, %d failed\n", successful, failed)
}
