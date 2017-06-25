package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcccc"))
}

func lengthOfLongestSubstring(s string) int {

	loc := make([]int, 128)

	for i := 0; i < 128; i++ {
		loc[i] = -1
	}
	begin, length := -1, 0
	for i, v := range s {
		if loc[v] > begin {
			begin = loc[v]
		}
		if i-begin > length {
			length = i - begin
		}
		loc[s[i]] = i
	}
	return length
}
