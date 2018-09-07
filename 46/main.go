package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	Mark := make(map[int]bool)
	for i, _ := range nums {
		Mark[i] = false
	}
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
	for i, _ := range nums {
		if mark[i] == false {
			mark[i] = true
			tmp = append(tmp, nums[i])
			Recurive(nums, mark, depth+1, tmp)
			mark[i] = false
			tmp = tmp[0 : len(tmp)-1]
		}
	}
}

func main() {
	ss := permute([]int{1, 2, 3, 4})
	for _, v := range ss {
		fmt.Println(v)
	}
}
