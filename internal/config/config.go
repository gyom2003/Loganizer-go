// Package config handles the reading and parsing of JSON configuration files
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// LogConfig represents a single log configuration entry
type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

// Config represents the complete configuration loaded from JSON
type Config []LogConfig

// ParseError represents an error that occurs during configuration parsing
type ParseError struct {
	File string
	Err  error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("failed to parse config file %s: %v", e.File, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

// LoadConfig reads and parses a JSON configuration file
func LoadConfig(configPath string) (Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, &ParseError{
			File: configPath,
			Err:  err,
		}
	}

	return config, nil
}
