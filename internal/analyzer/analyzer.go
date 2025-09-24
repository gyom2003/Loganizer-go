// Package analyzer handles log file analysis and provides custom error types
package analyzer

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/axellelanca/go_loganizer/internal/config"
)

// FileNotFoundError represents an error when a log file cannot be found or accessed
type FileNotFoundError struct {
	Path string
	Err  error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found or inaccessible: %s", e.Path)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

// AnalysisError represents an error that occurs during log analysis
type AnalysisError struct {
	LogID string
	Err   error
}

func (e *AnalysisError) Error() string {
	return fmt.Sprintf("analysis failed for log %s: %v", e.LogID, e.Err)
}

func (e *AnalysisError) Unwrap() error {
	return e.Err
}

// Result represents the result of analyzing a single log file
type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

// Analyzer handles the analysis of multiple log files concurrently
type Analyzer struct {
	wg      sync.WaitGroup
	results chan Result
}

// NewAnalyzer creates a new Analyzer instance
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		results: make(chan Result, 100), // Buffered channel to prevent blocking
	}
}

// AnalyzeAll processes all log configurations concurrently
func (a *Analyzer) AnalyzeAll(configs config.Config) []Result {
	// Start a goroutine for each log config
	for _, cfg := range configs {
		a.wg.Add(1)
		go a.analyzeLog(cfg)
	}

	// Start a goroutine to close the results channel when all analyses are done
	go func() {
		a.wg.Wait()
		close(a.results)
	}()

	// Collect all results
	var allResults []Result
	for result := range a.results {
		allResults = append(allResults, result)
	}

	return allResults
}

// analyzeLog analyzes a single log file
func (a *Analyzer) analyzeLog(cfg config.LogConfig) {
	defer a.wg.Done()

	result := Result{
		LogID:    cfg.ID,
		FilePath: cfg.Path,
	}

	// Check if file exists and is readable
	if _, err := os.Stat(cfg.Path); err != nil {
		if os.IsNotExist(err) {
			fileErr := &FileNotFoundError{
				Path: cfg.Path,
				Err:  err,
			}
			result.Status = "FAILED"
			result.Message = "File not found."
			result.ErrorDetails = fileErr.Error()
		} else {
			result.Status = "FAILED"
			result.Message = "File access error."
			result.ErrorDetails = err.Error()
		}
		a.results <- result
		return
	}

	// Check if file is readable
	file, err := os.Open(cfg.Path)
	if err != nil {
		fileErr := &FileNotFoundError{
			Path: cfg.Path,
			Err:  err,
		}
		result.Status = "FAILED"
		result.Message = "File not accessible."
		result.ErrorDetails = fileErr.Error()
		a.results <- result
		return
	}
	file.Close()

	// Simulate analysis with random sleep (50-200ms)
	sleepDuration := time.Duration(50+rand.Intn(151)) * time.Millisecond
	time.Sleep(sleepDuration)

	// If we reach here, analysis was successful
	result.Status = "OK"
	result.Message = "Analysis completed successfully."
	result.ErrorDetails = ""

	a.results <- result
}
