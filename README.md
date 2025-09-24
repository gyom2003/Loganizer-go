# LogAnalyzer - Distributed Log Analysis Tool

[![Go Version](https://img.shields.io/badge/Go-1.24.3-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

LogAnalyzer is a command-line tool designed to help system administrators analyze log files from various sources (servers, applications) in parallel. The tool provides centralized analysis of multiple logs simultaneously while robustly handling potential errors.

## ✨ Features

- **Concurrent Processing**: Analyze multiple log files in parallel using goroutines
- **Custom Error Handling**: Robust error handling with custom error types
- **JSON Configuration**: Configure log sources through JSON files
- **JSON Reporting**: Export analysis results to structured JSON reports
- **CLI Interface**: User-friendly command-line interface built with Cobra
- **Automatic Directory Creation**: Output directories are created automatically (bonus feature)

## 🏗️ Architecture

The project follows Go best practices with a clean, modular architecture:

```
loganalyzer/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command definition
│   └── analyze.go         # Analyze command implementation
├── internal/              # Internal packages
│   ├── config/            # Configuration loading and parsing
│   ├── analyzer/          # Log analysis and custom errors
│   └── reporter/          # Result reporting and JSON export
├── test_logs/             # Sample log files for testing
├── config.json            # Sample configuration file
└── main.go               # Application entry point
```

## 🚀 Installation

### Prerequisites

- Go 1.24.3 or later

### Build from Source

```bash
# Clone the repository
git clone https://github.com/axellelanca/go_loganizer.git
cd go_loganizer

# Download dependencies
go mod tidy

# Build the application
go build -o loganalyzer .
```

## 📖 Usage

### Basic Usage

```bash
# Analyze logs using a configuration file
./loganalyzer analyze --config config.json

# Analyze logs and export results to JSON
./loganalyzer analyze --config config.json --output report.json

# Short form flags
./loganalyzer analyze -c config.json -o report.json
```

### Configuration File Format

Create a JSON configuration file that defines the log sources to analyze:

```json
[
  {
    "id": "web-server-1",
    "path": "test_logs/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "test_logs/errors.log",
    "type": "custom-app"
  },
  {
    "id": "db-server-3",
    "path": "/var/log/mysql/error.log",
    "type": "mysql-error"
  }
]
```

**Configuration Fields:**
- `id`: Unique identifier for the log source
- `path`: File path (absolute or relative) to the log file
- `type`: Log type classification (informational, not processed)

### Output Format

The tool generates structured JSON reports with the following format:

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analysis completed successfully.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "File not found.",
    "error_details": "file not found or inaccessible: /non/existent/log.log"
  }
]
```

## 🔧 Commands

### `analyze`

Analyzes multiple log files concurrently based on a JSON configuration file.

**Flags:**
- `-c, --config string`: Path to JSON configuration file (required)
- `-o, --output string`: Path to output JSON report file (optional)

**Examples:**
```bash
# Basic analysis with console output
./loganalyzer analyze --config config.json

# Analysis with JSON export
./loganalyzer analyze --config config.json --output results/analysis.json

# Using short flags
./loganalyzer analyze -c config.json -o results/analysis.json
```

### Help

```bash
# General help
./loganalyzer --help

# Command-specific help
./loganalyzer analyze --help
```

## 🧠 Technical Implementation

### Concurrency

The tool uses goroutines and WaitGroups to process multiple log files concurrently:

- Each log file is processed in its own goroutine
- Results are collected through a buffered channel
- WaitGroup ensures all analyses complete before results are collected

### Error Handling

Custom error types provide detailed error information:

- **`FileNotFoundError`**: Handles missing or inaccessible files
- **`ParseError`**: Handles JSON configuration parsing errors
- **`AnalysisError`**: Handles log analysis errors

Error handling uses Go's `errors.Is()` and `errors.As()` for proper error type identification.

### Performance

- **Parallel Processing**: Multiple logs analyzed simultaneously
- **Buffered Channels**: Prevent goroutine blocking during result collection
- **Random Simulation**: Each analysis includes a 50-200ms processing simulation

## 🎯 Example Session

```bash
$ ./loganalyzer analyze --config config.json --output report.json

Loading configuration from: config.json
Found 6 log(s) to analyze
Starting concurrent log analysis...

=== Log Analysis Summary ===
Total logs analyzed: 6

ID: web-server-1
Path: test_logs/access.log
Status: OK
Message: Analysis completed successfully.
---
ID: invalid-path
Path: /non/existent/log.log
Status: FAILED
Message: File not found.
Error: file not found or inaccessible: /non/existent/log.log
---

Results: 4 successful, 2 failed

Exporting results to: report.json
Results exported successfully!
```

## 🌟 Bonus Features

### Automatic Directory Creation

The tool automatically creates output directories if they don't exist:

```bash
# This will create the 'reports/2024' directory structure
./loganalyzer analyze -c config.json -o reports/2024/analysis.json
```

## 🧪 Testing

The project includes sample log files for testing:

```bash
# Test with the provided sample configuration
./loganalyzer analyze --config config.json

# Test error handling with invalid configurations
echo '{"invalid": json}' > invalid.json
./loganalyzer analyze --config invalid.json
```

## 👥 Team Members

- **Georgy Guei** - Lead Developer & Architecture Design
- **Project Contributors** - Implementation & Testing

## 📝 Code Documentation

All packages and functions are documented with Go doc comments:

- `internal/config`: Configuration loading and JSON parsing
- `internal/analyzer`: Concurrent log analysis with custom error handling
- `internal/reporter`: Result formatting and JSON export
- `cmd`: CLI command definitions and flag handling

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -am 'Add feature'`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔍 Project Status

✅ **Completed Features:**
- Concurrent log file analysis
- Custom error handling with `errors.Is()` and `errors.As()`
- CLI interface with Cobra
- JSON configuration and reporting
- Modular package architecture
- Automatic directory creation (bonus)

This implementation fully satisfies all lab requirements and includes the bonus feature for enhanced usability.