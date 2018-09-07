package main

import "fmt"

func trap(height []int) int {

	left := 0
	left_height := height[0]
	result := 0
	for i := left; i < len(height); i++ {
		if height[i] >= left_height {
			left = i
			left_height = height[i]
		} else {
			result += left_height - height[i]
		}
		fmt.Println(i, result)
	}
	return result
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
