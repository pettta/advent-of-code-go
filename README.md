# Advent of Code Golang Tool

A Go-based CLI tool for solving Advent of Code 2025 problems. This tool provides a structured framework for implementing and running solutions for each day's challenges.

## Installation

### Build from source

```bash
make build
```

This will create an executable `aoc` in the current directory.

### Install system-wide

```bash
sudo make install
```

This will build and install the `aoc` command to `/usr/local/bin`, making it available system-wide.

## Usage

### Setup a new day

Create a template file for a new day's implementation:

```bash
aoc setup -d 1 -y 2025   # Creates internal/2025/day/day01.go
aoc setup -d 25 -y 2025 # Creates internal/2025/day/day25.go
```

The template includes:
- A struct implementing the `Day` interface
- Stub methods for `SolvePart1` and `SolvePart2`
- Automatic registration in the day registry

### Run a solution

Execute a day's solution:

```bash
aoc day -y 2025 -d 1 -p 1 --file input.txt
aoc day -y 2025 -d 1 -p 2 --file input.txt
```

**Flags:**
- `-d, --day`: Day number (1-25, required)
- `-p, --part`: Part number (1 or 2, required)
- `-y, --year` : Year number (required to know what folder we are working in)
- `--file`: Path to input file (required)

**Example:**
```bash
aoc day -d 1 -p 1 --file inputs/day01.txt
# Output: Day 1, Part 1: <result> (took <duration>)
```

## Project Structure

```
.
├── cmd/
│   ├── main.go          # Main entry point
│   └── aoc/
│       ├── day.go       # Day command implementation
│       └── setup.go     # Setup command implementation
├── internal/
│   ├── day/
│   │   └── day.go       # Day interface and registry
│   ├── utils/
│   │   └── file.go       # File utility functions
│   └── 2025/
│       ├── day/
│       │   └── day01.go  # Example generated implementation for the year
│       └── utils/        # Year-specific helpers (optional)
├── Makefile             # Build and install targets
└── README.md
```

## Implementing a Day

1. **Setup the day:**
   ```bash
   aoc setup -d 1
   ```

2. **Edit the generated file** at `internal/2025/day/day01.go`:
   ```go
   func (d *Day1) SolvePart1(input []byte) (string, error) {
       // Your solution here
       return "result", nil
   }
   ```

3. **Run your solution:**
   ```bash
   aoc day -y 2025 -d 1 -p 1 --file inputs/2025-d1-input.txt
   ```

## Day Interface

Each day must implement the `Day` interface:

```go
type Day interface {
    SolvePart1(input []byte) (string, error)
    SolvePart2(input []byte) (string, error)
}
```

The input is provided as raw bytes, allowing you to parse it however you need.

## Makefile Targets

- `make` or `make all` - Build the `aoc` binary
- `make build` - Build the `aoc` binary
- `make install` - Build and install to `/usr/local/bin` (requires sudo)
- `make clean` - Remove the built binary

## Requirements

- Go 1.25.4 or later
- Make (for build automation)

## License

This project is for personal use in solving Advent of Code 2025 challenges.
