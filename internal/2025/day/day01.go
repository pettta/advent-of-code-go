package day

import (
	"fmt"
	"strconv"
	daypkg "advent-of-code-go/internal/day"
	utils "advent-of-code-go/internal/utils"
)

type Day1 struct{}

func init() {
	daypkg.Days.RegisterDay(2025, 1, &Day1{})
}

func (d *Day1) SolvePart1(input []byte) (string, error) {
	lines := utils.ReadLines(input)
	dialLoc := 50 
	soln := 0 
	for _, line := range lines {
		if line == "" {
			continue
		}
		direction := line[0]
		steps, _ := strconv.Atoi(line[1:]) 
		if direction == 'R' { 
			dialLoc += steps 
		} else if direction == 'L' {
			dialLoc -= steps
		}
		dialLoc = (dialLoc % 100 + 100) % 100 // stupid wrong modulo defn in go
		if dialLoc == 0 {
			soln += 1 
		}
	}
	return fmt.Sprintf("%d", soln), nil
}

func (d *Day1) SolvePart2(input []byte) (string, error) {
	lines := utils.ReadLines(input)
	dialLoc := 50  
	prevLoc := 50 
	soln := 0 
	abs :=  func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for _, line := range lines {
		if line == "" {
			continue
		}
		direction := line[0]
		steps, _ := strconv.Atoi(line[1:])
		
		if steps > 100 {
			soln += abs(steps / 100)
			steps %= 100 
		} 
		
		if direction == 'R' { 
			dialLoc += steps 
		} else if direction == 'L' {
			dialLoc -= steps
		}
		
		if (dialLoc < 0 || dialLoc > 99){
			dialLoc = (dialLoc % 100 + 100) % 100 // stupid wrong modulo defn in go
			if (prevLoc != 0 && dialLoc != 0){
				soln += 1 
			}
		}
		
		if dialLoc == 0 {
			soln += 1 
		}
		
		prevLoc = dialLoc 
		
	}
	return fmt.Sprintf("%d", soln), nil
}
