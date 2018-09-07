package main

import "fmt"

func solveSudoku(board [][]byte) {
	solveSudokuDFS(board, 0, 0)

}
func solveSudokuDFS(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	if j >= 9 {
		return solveSudokuDFS(board, i+1, 0)
	}
	if board[i][j] == '.' {
		for k := 1; k <= 9; k++ {

			board[i][j] = '0' + byte(k)
			if check_valid(board, i, j) {
				if solveSudokuDFS(board, i, j+1) {
					return true
				}
			}
			board[i][j] = '.'
		}
	} else {
		return solveSudokuDFS(board, i, j+1)
	}
	return false
}

func check_valid(board [][]byte, x, y int) bool {
	var row, col, xx = make(map[byte]bool), make(map[byte]bool), make(map[byte]bool)
	i := 0
	//defer fmt.Println("xsdsd",i,row)
	for i = 0; i < len(board); i++ {
		r := board[x][i]
		if r != '.' && row[r] == true {
			return false
		} else {
			row[r] = true
		}
		c := board[i][y]
		if c != '.' && col[c] == true {
			return false
		} else {
			col[c] = true
		}
		xi := x/3*3 + i/3
		yi := y/3*3 + i%3
		val := board[xi][yi]
		if val != '.' && xx[val] == true {
			return false
		} else {
			xx[val] = true
		}
	}
	return true
}

/*
	xi = x/3*3 + i/3
	yi = y/3*3 + i%3

*/

func main() {
	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		board[i] = make([]byte, 9)
	}
	board[0] = []byte{'.', '.', '9', '7', '4', '8', '.', '.', '.'}
	board[1] = []byte{'7', '.', '.', '.', '.', '.', '.', '.', '.'}
	board[2] = []byte{'.', '2', '.', '1', '.', '9', '.', '.', '.'}
	board[3] = []byte{'.', '.', '7', '.', '.', '.', '2', '4', '.'}
	board[4] = []byte{'.', '6', '4', '.', '1', '.', '5', '9', '.'}
	board[5] = []byte{'.', '9', '8', '.', '.', '.', '3', '.', '.'}
	board[7] = []byte{'.', '.', '.', '8', '.', '3', '.', '2', '.'}
	board[6] = []byte{'.', '.', '.', '.', '.', '.', '.', '.', '6'}
	board[8] = []byte{'.', '.', '.', '2', '7', '5', '9', '.', '.'}

	// ["..9748...","7........",".2.1.9...","..7...24.",".64.1.59.",".98...3..","...8.3.2.","........6","...2759.."]

	solveSudoku(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Print("\n")
	}

}
