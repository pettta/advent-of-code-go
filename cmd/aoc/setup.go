package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	yearsgen "advent-of-code-go/internal/yearsgen"

	"github.com/spf13/cobra"
)

var (
	setupDay  int
	setupYear int
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup a new day",
	Long:  `Setup creates a template file for a new day's implementation`,
	Run: func(cmd *cobra.Command, args []string) {

		// Validate day & year inputs for folder creation etc .
		if setupDay < 1 || setupDay > 25 {
			fmt.Fprintf(os.Stderr, "Day must be between 1 and 25, got %d\n", setupDay)
			os.Exit(1)
		}
		if setupYear < 1 {
			fmt.Fprintf(os.Stderr, "Year must be a positive integer, got %d\n", setupYear)
			os.Exit(1)
		}
		yearFolderName := fmt.Sprintf("%d", setupYear)
		dayFileName := fmt.Sprintf("day%02d.go", setupDay)
		dayFilePath := filepath.Join("internal", yearFolderName, "day", dayFileName)
		inputFilePath := filepath.Join("inputs", yearFolderName, fmt.Sprintf("day%02d.txt", setupDay))
		if _, err := os.Stat(dayFilePath); err == nil {
			fmt.Fprintf(os.Stderr, "Day %d already exists at %s\n", setupDay, dayFilePath)
			os.Exit(1)
		}
		if err := os.MkdirAll(filepath.Dir(dayFilePath), 0o755); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create directories: %v\n", err)
			os.Exit(1)
		}
		if err := os.MkdirAll(filepath.Dir(inputFilePath), 0o755); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create input directories: %v\n", err)
			os.Exit(1)
		}

		// Create the template
		tmpl := `package day

import (
	"fmt"
	daypkg "advent-of-code-go/internal/day"
	utils "advent-of-code-go/internal/utils"
)

type Day{{.DayNum}} struct{}

func init() {
	daypkg.Days.RegisterDay({{.Year}}, {{.DayNum}}, &Day{{.DayNum}}{})
}

func (d *Day{{.DayNum}}) SolvePart1(input []byte) (string, error) {
	// TODO: Implement part 1
	return "", fmt.Errorf("not implemented")
}

func (d *Day{{.DayNum}}) SolvePart2(input []byte) (string, error) {
	// TODO: Implement part 2
	return "", fmt.Errorf("not implemented")
}
`

		// Verify template and file creation works
		t, err := template.New("day").Parse(tmpl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
			os.Exit(1)
		}
		file, err := os.Create(dayFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		data := struct {
			Year   int
			DayNum int
		}{
			Year:   setupYear,
			DayNum: setupDay,
		}
		if err := t.Execute(file, data); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created year %d day %d template at %s\n", setupYear, setupDay, dayFilePath)

		if err := yearsgen.Generate("."); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update year registry: %v\n", err)
			os.Exit(1)
		}

		// Curl the input file from Advent of Code's website
		url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", setupYear, setupDay)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
			os.Exit(1)
		}
		session := os.Getenv("AOC")
		if session != "" {
			req.AddCookie(&http.Cookie{Name: "session", Value: session})
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating input file: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "Failed to retrieve input file: %s\n Go grab it manually", resp.Status)
			os.Exit(1)
		}
		inputFile, err := os.Create(inputFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating input file: %v\n", err)
			os.Exit(1)
		}
		defer inputFile.Close()
		if _, err := io.Copy(inputFile, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing input file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Content from %s written to %s\n", url, inputFilePath)
	},
}

func init() {
	SetupCmd.Flags().IntVarP(&setupDay, "day", "d", 0, "Day number to setup")
	SetupCmd.Flags().IntVarP(&setupYear, "year", "y", 0, "Year number to setup")
	SetupCmd.MarkFlagRequired("day")
	SetupCmd.MarkFlagRequired("year")
}

// aggregator file generation handled by yearsgen
