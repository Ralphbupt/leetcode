package main

import "fmt"

func main() {

	uids := []int{1, 2, 4}
	for _, u := range uids {
		if u != 1 {
			uids = append(uids, u)
		}
	}
	fmt.Println(uids)
	x := 123333324

	fmt.Println(string(x))

	fmt.Printf("%v \n", string(x))
	fmt.Println(122)
}

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	max, tmp := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > 0 {
			tmp += nums[i]
		} else {
			if tmp < nums[i] {
				tmp = nums[i]
			}
		}
	}
	return max

}

/**/

/*
	max [i+1] = max[i] + nums[i]        nums[i] > 0
	max [i+i]  = max(max[i], nums[i])


*/
