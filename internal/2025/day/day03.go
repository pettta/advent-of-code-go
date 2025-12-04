package day

import (
	"fmt"
	"strconv"
	daypkg "advent-of-code-go/internal/day"
	utils "advent-of-code-go/internal/utils"
)

type Day3 struct{}

func init() {
	daypkg.Days.RegisterDay(2025, 3, &Day3{})
}

// thought this was a textbook heap problem, beacause i missed
// that you can only pick from after the first element for the second, etc... 
// went with a stack to solve this

type Stack struct {
	items []int
}
func (s *Stack) Push(i int) {s.items = append(s.items, i)}
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func findBankJoltage(line string, batteries int) string {
	s := &Stack{}
	n := len(line) 
	for loc, val := range line {
		currVal, _ := strconv.Atoi(string(val))
		//figure out best number for spot, s.t. min spots needed left ensured 
		for len(s.items) > 0 && currVal > s.Peek() && (n-loc) > (batteries - len(s.items)) {
			s.Pop()
		}
		if len(s.items) < batteries {
			s.Push(currVal)
		}
	}
	result := ""
	for _, v := range s.items {
		result += strconv.Itoa(v)
	}
	return result
}

func (d *Day3) SolvePart1(input []byte) (string, error) {
	soln := 0 
	lines := utils.ReadLines(input)
	for _, line := range lines {
		out := findBankJoltage(line, 2)
		val, _ := strconv.Atoi(out)
		soln += val
	}
	return fmt.Sprintf("%d", soln), nil
}


func (d *Day3) SolvePart2(input []byte) (string, error) {
	soln := 0 
	lines := utils.ReadLines(input)
	for _, line := range lines {
		out := findBankJoltage(line, 12)
		val, _ := strconv.Atoi(out)
		soln += val
	}
	return fmt.Sprintf("%d", soln), nil
}
