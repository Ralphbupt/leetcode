package main

import "fmt"

func searchInsert(nums []int, target int) int {

	// 0  1 2 3

	idx := 0

	for idx < len(nums) {
		if nums[idx] >= target {
			return idx
		}
		idx++
	}

	return idx

}

func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 4}, 6))
}
