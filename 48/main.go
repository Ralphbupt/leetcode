package main

import "fmt"

func rotate(matrix [][]int) {
	step := len(matrix) / 2
	width := len(matrix)

	for col := 0; col < len(matrix); col++ {
		for i := 0; i < step; i++ {
			matrix[i][col], matrix[width-1-i][col] = matrix[width-1-i][col], matrix[i][col]
		}
	}

	for row := 0; row < len(matrix); row++ {
		for i := 0; i <= row; i++ {
			//fmt.Println(i, row,matrix[row][i], matrix[i][row])
			matrix[row][i], matrix[i][row] = matrix[i][row], matrix[row][i]
		}
	}

}

func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
func main() {

	var matrix [][]int
	matrix = make([][]int, 0)
	matrix = append(matrix, []int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})

	rotate(matrix)
}
