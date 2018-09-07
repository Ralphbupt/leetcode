package main

import "fmt"

func canJump(nums []int) bool {
	max := 0
	for index, value := range nums {
		if value > 0 {
			if index+value > max {
				max = index + value
			}
		} else {
			if max == index {
				if index == len(nums)-1 { // 为什么加了这一步就好了呢
					return true
				}
				return false
			}
		}

	}
	return max >= len(nums)-1
}

func main() {
	nums := []int{3, 2, 1, 2, 4}
	fmt.Println(canJump(nums))

	nums = []int{3, 2, 1, 0, 3}

	fmt.Println(canJump(nums))

	nums = []int{0}

	fmt.Println(canJump(nums))
}
