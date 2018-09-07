package main

import (
	"fmt"
	"time"
)

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	if len(matrix) == 0 {
		return nil
	}
	left, up := 0, 0
	right, down := len(matrix[0])-1, len(matrix)-1
	for left <= right && up <= down {
		// from left to right
		for j := left; j <= right; j++ {
			result = append(result, matrix[up][j])
		}
		up++

		for i := up; i <= down; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if up <= down {
			for j := right; j >= left; j-- {
				result = append(result, matrix[down][j])
			}

		}
		down--

		if left <= right { // 最重要的一部， 可以不用判断奇偶
			for i := down; i >= up; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
	}
	return result
}

func TestI(row, col int) {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, 0)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			matrix[i] = append(matrix[i], i*col+j+1)
			fmt.Printf("%3d", i*col+j+1)
		}
		fmt.Println()
	}
	result := spiralOrder(matrix)
	fmt.Println(result)
	fmt.Println()
}

func main() {
	curWeek := (time.Now().YearDay() - 2) / 7

	fmt.Println(curWeek)


	now := time.Now()

	fmt.Println((now.Unix() / 86400), ((now.Unix()/86400 + 3) / 7), now.Weekday(), (now.YearDay()-2)/7, (now.YearDay()+3)/7)
	d, _ := time.ParseDuration("-24h")

	for i := 0; i < 100; i++ {

		now = now.Add(d)

		fmt.Println((now.Unix() / 86400), ((now.Unix()/86400 + 3) / 7), now.Weekday(), (now.YearDay()-2)/7, (now.YearDay()+3)/7)
	}

	/*
		TestI(1, 1)
		TestI(1, 2)
		TestI(1, 3)
		TestI(2, 3)
		TestI(3, 3)
		TestI(3, 4)
		TestI(3, 5)
		TestI(5, 5)
	*/
}
