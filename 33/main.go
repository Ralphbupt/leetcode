package main

import "fmt"

func main() {
	//	fmt.Println(search([]int{5,1,3},3))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6}, 2))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		//	fmt.Println(nums[mid] , nums[right], nums[left])

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
		//fmt.Println(left,right, mid)
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
