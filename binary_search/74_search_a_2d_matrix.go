package binarysearch

func SearchMatrix(matrix [][]int, target int) bool {
	numRows := len(matrix)
	numCols := len(matrix[0])

	left, right := 0, numRows-1

	for left < right {
		mid := (right + left) / 2
		lastNumInRow := matrix[mid][numCols-1]

		if lastNumInRow == target {
			return true
		}

		if lastNumInRow > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	currentRow := right

	left, right = 0, numCols-1
	for left <= right {
		mid := (right + left) / 2
		num := matrix[currentRow][mid]

		if num == target {
			return true
		}

		if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return false
}
