package graphs

func MaxAreaOfIsland(grid [][]int) int {
	numRows := len(grid)
	numCols := len(grid[0])
	maxArea := 0

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 || row == numRows || col < 0 || col == numCols || grid[row][col] == 0 {
			return 0
		}
		grid[row][col] = 0

		area := 1
		area += dfs(row+1, col)
		area += dfs(row-1, col)
		area += dfs(row, col+1)
		area += dfs(row, col-1)

		return area
	}

	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == 1 {
				area := dfs(row, col)
				maxArea = max(area, maxArea)
			}
		}
	}

	return maxArea
}
