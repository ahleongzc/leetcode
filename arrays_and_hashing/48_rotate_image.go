package arraysandhashing

func rotate(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < row; col++ {
			tmp := matrix[row][col]
			matrix[row][col] = matrix[col][row]
			matrix[col][row] = tmp
		}
	}

	for row := 0; row < len(matrix); row++ {
		left, right := 0, len(matrix[0])-1
		for left <= right {
			tmp := matrix[row][left]
			matrix[row][left] = matrix[row][right]
			matrix[row][right] = tmp
			left++
			right--
		}
	}

}
