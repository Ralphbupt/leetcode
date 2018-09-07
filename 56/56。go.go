package main

import (
	"math"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type SortInteval []Interval

func (s SortInteval) Len() int {
	return len(s)
}

func (s SortInteval) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func (s SortInteval) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 0 {
		return intervals
	}
	sort.Sort(SortInteval(intervals))
	result := make([]Interval, 0)

	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1].End < intervals[i].Start {
			result = append(result, intervals[i])
		} else {
			if result[len(result)-1].End < intervals[i].End {
				result[len(result)-1].End = intervals[i].End
			}

		}
	}

	return result
}

/*
	res.push_back(ins[0]);
    for (int i = 1; i < ins.size(); i++) {
        if (res.back().end < ins[i].start) res.push_back(ins[i]);
        else
            res.back().end = max(res.back().end, ins[i].end);
    }
    ret
*/

/*

Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].

*/
func main() {

}
