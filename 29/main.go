package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	pos := divisor > 0 && dividend > 0 || divisor < 0 && dividend < 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := 0
	p := 1
	for {
		if (p<<1)*divisor > dividend {
			break
		}
		p = p << 1
	}
	for dividend > 0 && p > 0 {
		if dividend >= divisor*p {
			res += p
			dividend -= divisor * p
		}
		p = p >> 1
	}

	if pos {
		return res
	} else {
		return -res
	}
}

func main() {
	test(1, 1, 1)
	test(math.MaxInt32, 1, math.MaxInt32)
	test(12, 3, 4)
	test(-100, 4, -25)
	test(math.MinInt32, -1, math.MaxInt32)
	test(0, 0, math.MaxInt32)
	test(2, 5, 0)
	test(3, 5, 0)
	test(4, 5, 0)
	test(5, 5, 1)
	test(6, 5, 1)
	test(80, 5, 16)
	fmt.Println("if no error occured , passed")
}

func test(dividend int, divisor int, expected int) {
	if ret := divide(dividend, divisor); ret != expected {
   		fmt.Println(str)
	}
}
