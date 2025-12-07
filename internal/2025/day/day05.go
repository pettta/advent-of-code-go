package day

import (
	daypkg "advent-of-code-go/internal/day"
	utils "advent-of-code-go/internal/utils"
	"fmt"
)

type Day5 struct{}

func init() {
	daypkg.Days.RegisterDay(2025, 5, &Day5{})
}

func (d *Day5) SolvePart1(input []byte) (string, error) {
	lines := utils.ReadLines(input)
	fmt.Println("lines=", lines)
	return "", fmt.Errorf("not implemented")
}

func (d *Day5) SolvePart2(input []byte) (string, error) {
	// TODO: Implement part 2
	return "", fmt.Errorf("not implemented")
}
