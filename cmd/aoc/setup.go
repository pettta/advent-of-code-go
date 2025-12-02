package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	setupDay int
	setupYear int
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup a new day",
	Long:  `Setup creates a template file for a new day's implementation`,
	Run: func(cmd *cobra.Command, args []string) {
		
		
		// Validate day inputs (not bothering with year) 
		if setupDay < 1 || setupDay > 25 {
			fmt.Fprintf(os.Stderr, "Day must be between 1 and 25, got %d\n", setupDay)
			os.Exit(1)
		}
		yearFolderName := fmt.Sprintf("year%d", setupYear)
		dayFileName := fmt.Sprintf("day%02d.go", setupDay)
		dayFilePath := filepath.Join("internal", yearFolderName, "day", dayFileName)
		if _, err := os.Stat(dayFilePath); err == nil {
			fmt.Fprintf(os.Stderr, "Day %d already exists at %s\n", setupDay, dayFilePath)
			os.Exit(1)
		}

		
		// Create the template
		tmpl := `package day

import (
	"fmt"
)

type Day{{.DayNum}} struct{}

func init() {
	Days.RegisterDay({{.DayNum}}, &Day{{.DayNum}}{})
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
			DayNum int
		}{
			DayNum: setupDay,
		}
		if err := t.Execute(file, data); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
			os.Exit(1)
		}
		

		fmt.Printf("Created day %d template at %s\n", setupDay, dayFilePath)
	},
}

func init() {
	SetupCmd.Flags().IntVarP(&setupDay, "day", "d", 0, "Day number to setup")
	SetupCmd.Flags().IntVarP(&setupYear, "year", "y", 0, "Year number to setup")
	SetupCmd.MarkFlagRequired("day")
	SetupCmd.MarkFlagRequired("year")
}
