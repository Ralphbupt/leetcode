package main

import (
	"sort"
	"fmt"
)

var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res = make([][]int, 0)

	Recurcive(candidates, target, 0, make([]int, 0))
	return res
}

func Recurcive(candidates []int, target int, index int, tmp []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		xx := make([]int, len(tmp))
		copy(xx, tmp)
		res = append(res, xx)
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {	// this is important
			// because candidates is sorted , so this could be used to escapes the
			// sfjdifj
			continue
		}
		tmp = append(tmp, candidates[i])
		Recurcive(candidates, target-candidates[i], i+1,tmp)
		tmp = tmp[0:len(tmp)-1]
	}
}

func main() {
	res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	for _, v := range res {
		fmt.Println(v)
	}
}