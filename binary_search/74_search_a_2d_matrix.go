package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	numRows := len(matrix)
	numCols := len(matrix[0])

	left, right := 0, numRows-1
	for left < right {
		mid := (left + right) / 2
		if target > matrix[mid][numCols-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	targetRow := left
	left, right = 0, numCols-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[targetRow][mid] == target {
			return true
		}

		if target > matrix[targetRow][mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
