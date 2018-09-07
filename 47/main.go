package main

import (
	"fmt"
	"sort"
)

var res [][]int

func permuteUnique(nums []int) [][]int {
	Mark := make(map[int]bool)
	for i, _ := range nums {
		Mark[i] = false
	}
	sort.Ints(nums)
	depth := 0
	res = make([][]int, 0)
	Recurive(nums, Mark, depth, make([]int, 0))
	return res
}

func Recurive(nums []int, mark map[int]bool, depth int, tmp []int) {
	if depth == len(nums) {
		xx := make([]int, len(tmp))
		copy(xx, tmp)
		res = append(res, xx)
	}
	for i, v := range nums {
		if mark[i-1] && i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if mark[i] == false {
			mark[i] = true
			tmp = append(tmp, v)
			Recurive(nums, mark, depth+1, tmp)
			mark[i] = false
			tmp = tmp[0 : len(tmp)-1]
		}
	}
}

func main() {
	ss := permuteUnique([]int{1, 1, 0, 1})
	for _, v := range ss {
		fmt.Println(v)
	}
}
