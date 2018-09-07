package main

import "fmt"

func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		//fmt.p
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			swap(nums, i, nums[i]-1)
		}
		///fmt.Println(i,nums)
	}

	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//  1   2   3   4
//  -1  1   2   4

func main() {
	fmt.Println(firstMissingPositive([]int{-1, 4, 2, 1, 9, 10}))
	fmt.Println(firstMissingPositive([]int{1, 2, 3}))
	fmt.Println(firstMissingPositive([]int{0}))
	fmt.Println(firstMissingPositive([]int{}))
	fmt.Println(firstMissingPositive([]int{0, 1, 2, 3}))
	fmt.Println(firstMissingPositive([]int{5, 2, 4}))
	fmt.Println(firstMissingPositive([]int{-1, 1, 2, 4}))
	fmt.Println(firstMissingPositive([]int{-1, 5, 2, 4}))
	fmt.Println(firstMissingPositive([]int{-1, 9, 2, 1, 3, 4}))

}
