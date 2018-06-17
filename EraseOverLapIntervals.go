package main

import "sort"
import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func main() {
	intervals := []Interval{{1, 2}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals []Interval) int {
	length := len(intervals)
	if length == 0 {
		return 0
	}
	sort.Sort(Is(intervals))
	cnt := 1
	end := intervals[0].End
	for i := 1; i < length; i++ {
		if intervals[i].Start < end {
			continue
		}
		end = intervals[i].End
		cnt++
	}
	return len(intervals) - cnt

}

type Is []Interval

func (is Is) Len() int           { return len(is) }
func (is Is) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is Is) Less(i, j int) bool { return is[i].End < is[j].End }
