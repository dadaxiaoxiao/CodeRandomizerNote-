package leetcode0059

// generateMatrix
// leetcode59 螺旋矩阵 II
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	top, left, bottom, right := 0, 0, n-1, n-1
	for left <= right && top <= bottom {
		// top  left > right
		for column := left; column <= right; column++ {
			matrix[top][column] = num
			num++
		}
		top++

		// right top > bottom
		for row := top; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		right--

		if left <= right && top <= bottom {
			// bottom right > left
			for column := right; column >= left; column-- {
				matrix[bottom][column] = num
				num++
			}
			bottom--

			// left  bottom > top
			for row := bottom; row >= top; row-- {
				matrix[row][left] = num
				num++
			}
			left++
		}
	}
	return matrix
}
