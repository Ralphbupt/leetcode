package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		t := target - v
		if s, ok := m[t]; ok {
			result := make([]int, 2)
			result[1] = s
			result[0] = i
			return result
		} else {
			m[v] = i
		}
	}
	return nil
}

// 没什么问题， 一次 AC
