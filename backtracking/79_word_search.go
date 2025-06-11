package backtracking

func Exist(board [][]byte, word string) bool {
	numRows := len(board)
	numCols := len(board[0])

	var dfs func(row, col int, index int) bool
	dfs = func(row, col int, index int) bool {
		if row < 0 || col < 0 || row == numRows || col == numCols {
			return false
		}

		if index == len(word) {
			return false
		}

		if board[row][col] == '$' {
			return false
		}

		if word[index] != board[row][col] {
			return false
		}

		if index == len(word)-1 {
			return true
		}

		currChar := board[row][col]
		board[row][col] = '$'

		result := dfs(row+1, col, index+1) || dfs(row-1, col, index+1) || dfs(row, col+1, index+1) || dfs(row, col-1, index+1)

		board[row][col] = currChar

		return result
	}

	for row := range numRows {
		for col := range numCols {
			if board[row][col] != word[0] {
				continue
			}

			if dfs(row, col, 0) {
				return true
			}
		}
	}

	return false
}
