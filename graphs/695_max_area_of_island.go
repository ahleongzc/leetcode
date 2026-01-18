package graphs

func maxAreaOfIsland(grid [][]int) int {
	numRows, numCols := len(grid), len(grid[0])
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	curr := 0
	output := 0

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row == numRows || col == numCols {
			return
		}

		if grid[row][col] == 0 {
			return
		}

		curr++
		grid[row][col] = 0
		output = max(curr, output)

		for _, direction := range directions {
			dx, dy := direction[0], direction[1]
			dfs(row+dx, col+dy)
		}
	}

	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == 1 {
				curr = 0
				dfs(row, col)
			}
		}
	}

	return output
}
