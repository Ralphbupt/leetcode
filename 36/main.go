package main

func isValidSudoku(board [][]byte) bool {

	var row, col, xx map[byte]bool
	for i := 0; i < len(board); i++ {
		row = make(map[byte]bool)
		col = make(map[byte]bool)
		xx = make(map[byte]bool)
		for j := 0; j < len(board); j++ {
			r := board[i][j]
			if r != '.' && row[r] == true {
				return false
			} else {
				row[r] = true
			}
			c := board[j][i]
			if c != '.' && col[c] == true {
				return false
			} else {
				col[c] = true
			}
			x := (i*3 + j/3) % 9
			y := (i / 3 * 3) + j%3
			val := board[x][y]

			if val != '.' && xx[val] == true {
				return false
			} else {
				xx[val] = true
			}

		}
	}
	return true

}

/*

(i*3 + j/3)%9, (i/3 *3)+j%3

00 01 02	03 04 05 	06 07 08		00 01 02	30 31 32	60 61 62
10 11 12	13 14 15	16 17 18		03 04 05	33 34 35	63 64 65
20 21 22	23 24 25	26 27 28		06 07 08	36 37 38	66 67 68

30 31 32    33 34 35 	36 37 38		10 11 12	40 41 42	70 71 72
40 41 42	43 44 45	46 47 48		13 14 15	43 44 45	73 74 75
50 51 52	53 54 55 	56 57 58		16 17 18	46 47 48	76 77 78


60 61 62	63 64 65 	66 67 78 		20 21 22	50 51 52	80 81 82
70 71 72	73 74 75 	76 77 78 		23 24 25	53 54 55	83 84 85
80 81 82    83 84 85 	86 87 88 		26 27 28	56 57 58	86 87 88

*/

func main() {

}
