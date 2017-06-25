package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	//  1 2 3 4 5 6
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if nums[mid] == target {
		l, r := mid, mid
		for l >= 0 && nums[l] == target {
			l--
		}
		for r < len(nums) && nums[r] == target {
			r++
		}

		return []int{l + 1, r - 1}
	} else {
		return []int{-1, -1}
	}
}

func main() {
	fmt.Print(searchRange([]int{8}, 8))
}
