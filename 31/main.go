package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 5, 4, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	length := len(nums) - 1

	pos := length

	for pos > 0 {
		if nums[pos-1] < nums[pos] {
			break
		}
		pos--
	}
	if pos != 0 {
		end := length
		for end > 0 {
			if nums[end] > nums[pos-1] {
				break
			}
			end--
		}
		swap(nums, pos-1, end)
	}
	sort.Ints(nums[pos : length+1])

}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func test() {

}
