package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	left, up := 0, 0
	right, down := n-1, n-1

	k := 1

	for left <= right && up <= down {
		// from left to right
		for j := left; j <= right; j++ {
			matrix[up][j] = k
			k = k + 1
		}
		up++

		for i := up; i <= down; i++ {
			matrix[i][right] = k
			k = k + 1
			matrix[i][right] = k
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
