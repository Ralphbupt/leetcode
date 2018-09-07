package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	answer := make([][]string, 0)
	board := make([][]byte, n)

	// init board
	for i := 0; i < n; i++ {
		slice := make([]byte, 0)
		for i := 0; i < n; i++ {
			slice = append(slice, '.')
		}
		board[i] = slice
	}

	solve(&answer, board, n, 0)

	return answer
}

func solve(result *[][]string, board [][]byte, n int, row int) {

	if n == row {
		re := make([]string, 0)
		for i := 0; i < n; i++ {
			str := string(board[i])
			re = append(re, str)
		}
		*result = append(*result, re)
		return
	}
	for i := 0; i < n; i++ { // recur
		if check(n, board, row, i) {
			board[row][i] = 'Q'
			solve(result, board, n, row+1)
			board[row][i] = '.'
		}
	}

}

func check(n int, board [][]byte, row, col int) bool {
	for i := 0; i < n; i++ {
		if board[row][i] == 'Q' || board[i][col] == 'Q' {
			return false
		}
	}
	for i := -n; i < n; i++ {
		r, c, c1 := row+i, col+i, col-i
		if r >= 0 && r < n && c >= 0 && c < n {
			if board[r][c] == 'Q' {
				return false
			}
		}
		if r >= 0 && r < n && c1 >= 0 && c1 < n {
			if board[r][c1] == 'Q' {
				return false
			}
		}
	}
	return true
}

func main() {
	res := solveNQueens(4)
	for _, v := range res {
		for _, v1 := range v {
			fmt.Println(v1)
		}
		fmt.Println()
	}
}
