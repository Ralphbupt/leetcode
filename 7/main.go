package main

import (
	"fmt"

	"math"
)

func main() {
	fmt.Println(reverse(12345678996))
}

func reverse(x int) int {
	if -10 < x && x < 10 {
		return x
	}
	pos := 1
	if x < 0 {
		pos = -1
		x = -x
	}
	result := 0
	for x > 0 {
		result = result*10 + x%10
		if result > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	return pos * result
}
