package main

import "fmt"

func main() {
	m := 3
	n := 1 & m
	fmt.Printf("%T %T\n", m, n)
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	result := 0
	for x > 0 || y > 0 {
		if x&1 != y&1 {
			result++
		}
		x, y = x>>1, y>>1
	}
	return result
}
