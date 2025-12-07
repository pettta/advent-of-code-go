package day

import (
	daypkg "advent-of-code-go/internal/day"
	utils "advent-of-code-go/internal/utils"
	"fmt"
)

type Day4 struct{}

func init() {
	daypkg.Days.RegisterDay(2025, 4, &Day4{})
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

// could have had part one by faster by not coupling the logic
// but idrc
func getRolls(lines []string) (int, []string) {
	soln := 0
	newLines := make([]string, len(lines))
	for row, line := range lines {
		for col := range line {
			// 64 = @ = paper roll
			// 46 = . = blank space
			TL := row > 0 && col > 0 && (col-1) < len(lines[row-1]) && lines[row-1][col-1] == 64
			TM := row > 0 && col < len(lines[row-1]) && lines[row-1][col] == 64
			TR := row > 0 && (col+1) < len(lines[row-1]) && lines[row-1][col+1] == 64
			ML := col > 0 && (col-1) < len(lines[row]) && lines[row][col-1] == 64
			MR := (col+1) < len(lines[row]) && lines[row][col+1] == 64
			BL := (row+1) < len(lines) && col > 0 && (col-1) < len(lines[row+1]) && lines[row+1][col-1] == 64
			BM := (row+1) < len(lines) && col < len(lines[row+1]) && lines[row+1][col] == 64
			BR := (row+1) < len(lines) && (col+1) < len(lines[row+1]) && lines[row+1][col+1] == 64
			// stupid that you cant cast bools to int in go lol
			sum := utils.B2I(TL) + utils.B2I(TM) + utils.B2I(TR) + utils.B2I(ML) + utils.B2I(MR) + utils.B2I(BL) + utils.B2I(BM) + utils.B2I(BR)
			if lines[row][col] == 64 {
				if sum < 4 {
					newLines[row] += "x"
					soln++
					continue
				}
			}
			newLines[row] += string(lines[row][col])
		}
	}
	return soln, newLines
}

func (d *Day4) SolvePart1(input []byte) (string, error) {
	lines := utils.ReadLines(input)
	value, _ := getRolls(lines)
	return fmt.Sprintf("%d", value), nil
}

func (d *Day4) SolvePart2(input []byte) (string, error) {
	lines := utils.ReadLines(input)
	soln := 0
	out := 999
	for out > 0 {
		var newLines []string
		out, newLines = getRolls(lines)
		// printLines(newLines)
		// fmt.Println("\nout=", out)
		soln += out
		lines = newLines
	}
	return fmt.Sprintf("%d", soln), nil
}
