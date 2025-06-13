package graphs

func Solve(board [][]byte) {
	numRows := len(board)
	numCols := len(board[0])

	var dfs func(row, col int)
	dfs = func(row int, col int) {
		if row < 0 || col < 0 || row == numRows || col == numCols {
			return
		}

		if board[row][col] == byte('?') || board[row][col] == byte('X') {
			return
		}

		board[row][col] = byte('?')

		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for row := range numRows {
		for col := range numCols {
			if row != 0 && row != numRows-1 && col != 0 && col != numCols-1 {
				continue
			}
			dfs(row, col)
		}
	}

	for row := range numRows {
		for col := range numCols {
			if board[row][col] == byte('?') {
				board[row][col] = byte('O')
			} else {
				board[row][col] = byte('X')
			}
		}
	}
}
