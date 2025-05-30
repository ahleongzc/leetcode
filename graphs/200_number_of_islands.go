package graphs

func NumIslands(grid [][]byte) int {
	num := 0
	num_rows := len(grid)
	num_cols := len(grid[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row == num_rows || col < 0 || col == num_cols || grid[row][col] == byte('0') {
			return
		}

		grid[row][col] = byte('0')
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for row := range num_rows {
		for col := range num_cols {
			if grid[row][col] == byte('1') {
				dfs(row, col)
				num += 1
			}
		}
	}

	return num
}
