package graphs

func FindMaxFish(grid [][]int) int {
	numRows := len(grid)
	numCols := len(grid[0])

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 || col < 0 || row == numRows || col == numCols {
			return 0
		}
		if grid[row][col] == 0 {
			return 0
		}

		total := grid[row][col]
		grid[row][col] = 0

		total += dfs(row+1, col)
		total += dfs(row-1, col)
		total += dfs(row, col+1)
		total += dfs(row, col-1)

		return total
	}

	fishes := 0

	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == 0 {
				fishes = max(fishes, dfs(row, col))
			}
		}
	}

	return fishes
}
