package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}
func findMinArrowShots(points [][]int) int {
	length := len(points)
	if length == 0 {
		return length
	}

	sort.Sort(IS(points))
	cnt := 1
	end := points[0][1]
	for i := 1; i < length; i++ {
		if points[i][0] <= end {
			continue
		}
		cnt++
		end = points[i][1]
	}
	return cnt
}

type IS [][]int

func (is IS) Len() int           { return len(is) }
func (is IS) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is IS) Less(i, j int) bool { return is[i][1] < is[j][1] }
