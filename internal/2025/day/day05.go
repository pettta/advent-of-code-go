package day

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
	daypkg "advent-of-code-go/internal/day"
	utils "advent-of-code-go/internal/utils"
	
)

type Day5 struct{}

func init() {
	daypkg.Days.RegisterDay(2025, 5, &Day5{})
}

type Interval struct {
	start int 
	end int 
}
func (iv Interval) ContainsPoint(p int) bool {
    return iv.start <= int(p) && int(p) <= iv.end
}


type Node struct {
	interval Interval 
	max int 
	left *Node 
	right *Node 
}
func insertNode(root *Node, iv Interval) *Node {
	if root == nil {
		return &Node{
			interval: iv,
			max: iv.end,
			left: nil,
			right: nil,
		}
	}
	if iv.start < root.interval.start {
		root.left = insertNode(root.left, iv)
	} else {
		root.right = insertNode(root.right, iv)
	}
	root.max = root.interval.end 
	if root.left != nil && root.left.max > root.max {
		root.max = root.left.max 
	}
	if root.right != nil && root.right.max > root.max {
		root.max = root.right.max 
	}
	return root 
}


type IntervalTree struct { 
	root *Node 
}
func (t *IntervalTree) FindAnyContaining(p int) (iv Interval, ok bool) {
	n := t.root
	for n != nil {
		if n.interval.ContainsPoint(p) {
			return n.interval, true 
		}
		if n.left != nil && n.left.max >= p {
			n = n.left 
		} else {
			n = n.right 
		}
	}
	return Interval{}, false
}
func (t *IntervalTree) insert(iv Interval) { 
	t.root = insertNode(t.root, iv)
}



func (d *Day5) SolvePart1(input []byte) (string, error) {
	var tree IntervalTree 
	soln := 0
	lines := utils.ReadLines(input)
	isReadMode := false 
	for _, line := range lines {
		if line == "" {
			if !isReadMode {
				isReadMode = true
			}
			continue
		}
		if !isReadMode { 
			endPoints := strings.SplitN(line, "-", 2)
			start, _ := strconv.Atoi(endPoints[0])
			end, _ := strconv.Atoi(endPoints[1])
			// fmt.Println("range:", start, "to", end)
			tree.insert(Interval{start: start, end: end})
		} else {
			searchPoint, _ := strconv.Atoi(line)
			if _, ok := tree.FindAnyContaining(searchPoint); ok {
				// fmt.Println("point", line, "is contained in interval", iv)
				soln += 1 
			} else {
				// fmt.Println("point", line, "is NOT contained in any interval")
			}
		}
	}
	return fmt.Sprintf("%d", soln), nil
}

// seemed pretty separate to part one lol 
func (d *Day5) SolvePart2(input []byte) (string, error) {
	lines := utils.ReadLines(input)
    var intervals []Interval
    for _, line := range lines {
        if line == "" {
            break  
        }
        endPoints := strings.SplitN(line, "-", 2)
        start, _ := strconv.Atoi(endPoints[0])
        end, _ := strconv.Atoi(endPoints[1])
        intervals = append(intervals, Interval{start: start, end: end})
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i].start != intervals[j].start {
            return intervals[i].start < intervals[j].start
        }
        return intervals[i].end < intervals[j].end
    })
    merged := []Interval{intervals[0]}
    for i := range intervals[1:] {
        current := intervals[i+1]
        lastMerged := &merged[len(merged)-1]
        if lastMerged.end >= current.start-1 {
            if current.end > lastMerged.end {
                lastMerged.end = current.end
            }
        } else {
            merged = append(merged, current)
        }
    }
    totalCount := 0
    for _, iv := range merged {
        totalCount += (iv.end - iv.start + 1)
    }
    return fmt.Sprintf("%d", totalCount), nil
}
