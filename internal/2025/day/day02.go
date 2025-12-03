package day

import (
	"fmt"
	"strings"
	"strconv"
	// "math"
	daypkg "advent-of-code-go/internal/day"
	utils "advent-of-code-go/internal/utils"
)

type Day2 struct{}

func init() {
	daypkg.Days.RegisterDay(2025, 2, &Day2{})
}

// was curious if not using a string-based approach would be loads faster even just whipping it up quickly (lots of float/int casts, doing multiple calls, etc)
// at least my approach wasnt (was about the same exact speed) but im dumb and i wont claim to be mr golang lol

// func (d *Day2) SolvePart1(input []byte) (string, error) {
// 	soln := 0 
// 	firstHalf := func(n int) int {
// 		d := int(math.Floor(math.Log10(float64(n)))) + 1
// 		h := d / 2 
// 		return int(math.Floor(float64(n) / math.Pow(10, float64(d-h)),	))
// 	}
// 	secondHalf := func(n int) int {
// 		d := int(math.Floor(math.Log10(float64(n)))) + 1
// 		h := d / 2
// 		return n % int(math.Pow(10, float64(d-h)))
// 	}
// 	intLen := func(n int) int {
// 		if n == 0 {
// 			return 1
// 		}
// 		return int(math.Floor(math.Log10(float64(n)))) + 1
// 	}
// 	for _, line := range utils.ReadCSV(input) { 
// 		if line == "" {
// 			continue 
// 		}
// 		endPoints := strings.SplitN(line, "-", 2)
// 		start, _ := strconv.Atoi(endPoints[0])
// 		end, _ := strconv.Atoi(endPoints[1])
// 		for val := range end-start+1 {
// 			int approach: 
// 			if (intLen(start+val)%2 == 0) && (firstHalf(start+val) == secondHalf(start+val)) {
// 				soln += (start+val)
// 			}	
// 		}
// 	}
// 	return fmt.Sprintf("%d", soln), nil
// }

func (d *Day2) SolvePart1(input []byte) (string, error) {
	soln := 0 	
	for _, line := range utils.ReadCSV(input) { 
		if line == "" {
			continue 
		}
		endPoints := strings.SplitN(line, "-", 2)
		start, _ := strconv.Atoi(endPoints[0])
		end, _ := strconv.Atoi(endPoints[1])
		
		for val := range end-start+1 {
			strVal := fmt.Sprintf("%d", start+val)
			len := len(strVal) 
			if (len%2==0) && (strVal[0:len/2] == strVal[len/2:len]) {
				soln += (start+val)
			}	
		}
		
	}
	return fmt.Sprintf("%d", soln), nil
}

func (d *Day2) SolvePart2(input []byte) (string, error) {
	soln := 0 
	for _, line := range utils.ReadCSV(input) {
		if line == "" {
			continue
		}
		endPoints := strings.SplitN(line, "-", 2)
		start, _ := strconv.Atoi(endPoints[0])
		end, _ := strconv.Atoi(endPoints[1])
		for val := range end-start+1 {
			strVal := fmt.Sprintf("%d", start+val)
			len := len(strVal) 
			// sliding window from size k=2...k=n/2, with n/k duplicates?  
			for k := 1; k <= len/2; k++ {
				// skip window sizes that dont divide n evenly 
				if len%k != 0 {
					continue
				}
				prevWindow := "" 
				isInvalid := true  
				for i := range len/k {
					// fmt.Println("prevwindow=", strVal[i*k:(i+1)*k], "full string=", strVal)
					if prevWindow != "" && prevWindow != strVal[i*k:(i+1)*k] {
						isInvalid = false 
						break 
					}
					prevWindow = strVal[i*k:(i+1)*k]
					
				}
				if isInvalid {
					// fmt.Println("adding string=", strVal)
					soln += (start+val)
					break 
				}
				
			}
			
		}
		// return "", fmt.Errorf("not implemented")
	}
	return fmt.Sprintf("%d", soln), nil
}
