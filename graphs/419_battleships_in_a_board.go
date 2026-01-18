package graphs

func countBattleships(board [][]byte) int {
	numRows, numCols := len(board), len(board[0])
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row == numRows || col == numCols {
			return
		}

		if board[row][col] == '.' {
			return
		}

		board[row][col] = '.'

		for _, direction := range directions {
			dx, dy := direction[0], direction[1]
			dfs(row+dx, col+dy)
		}
	}

	battleships := 0
	
	for row := range numRows {
		for col := range numCols {
			if board[row][col] == 'X' {
				battleships++
				dfs(row, col)
			}
		}
	}

	return battleships
}
