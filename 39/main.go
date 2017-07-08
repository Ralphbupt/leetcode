package main

import (
	"fmt"
	"sort"
)

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res = make([][]int, 0)
	Recurcive(candidates, 0, target, make([]int, 0))
	return res
}

func Recurcive(candidates []int, left int, target int, tmp []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		xx := make([]int, len(tmp))
		copy(xx, tmp)
		res = append(res, xx)
		return
	}
	for i := left; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		Recurcive(candidates, i, target-candidates[i], tmp)
		tmp = tmp[0 : len(tmp)-1]
	}
}

func main() {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	for _, v := range res {
		fmt.Println(v)
	}
}
