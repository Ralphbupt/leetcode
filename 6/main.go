package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("123456789", 4))

}
func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	length := len(s)
	if length <= 1 {
		return s
	}

	results := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		results[i] = make([]byte, 0)
	}

	for i, _ := range s {
		tmp := i % (2*numRows - 2)
		if tmp >= numRows {
			tmp = 2*numRows - tmp - 2
		}
		results[tmp] = append(results[tmp], s[i])
	}
	result := ""

	for _, s := range results {
		result += string(s)
	}
	return result
}

/*
Z   Y   Y
A X N S Q
Z   S   M
*/
