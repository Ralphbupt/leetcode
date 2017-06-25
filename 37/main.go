package main


func solveSudoku(board [][]byte){
	solveSudokuDFS(board,0,0)

}
func solveSudokuDFS(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	if j >= 9 {
		return solveSudokuDFS(board, i+1, 0)
	}
	if board[i][j] == '.' {
		for k := 1; k <= 9; i++ {
			board[i][j] = '0' + byte(k)
			if check_valid(board, i, j) && solveSudokuDFS(board, i, j+1) {
				return true
			}
			board[i][j] = '.'
		}
	} else {
		return solveSudokuDFS(board, i, j+1)
	}
	return false
}

func check_valid(board [][]byte, x, y int) bool {
	var row, col, xx map[byte]bool
	for i := 0; i < len(board); i++ {
		r := board[x][i]
		if r == '.' || row[r] == true {
			return false
		} else {
			row[r] = true
		}
		c := board[i][y]
		if c == '.' || col[c] == true {
			return false
		} else {
			col[c] = true
		}
		xi := x/3*3 + i/3
		yi := y/3*3 + i%3
		val := board[xi][yi]
		if val == '.' || xx[val] == true {
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

	solveSudoku(board)
}
