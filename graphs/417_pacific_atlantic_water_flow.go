package graphs

func PacificAtlantic(heights [][]int) [][]int {
	numRows := len(heights)
	numCols := len(heights[0])

	canReachFromPacific := make([][]bool, numRows)
	for i := range numRows {
		canReachFromPacific[i] = make([]bool, numCols)
	}

	canReachFromAtlantic := make([][]bool, numRows)
	for i := range numRows {
		canReachFromAtlantic[i] = make([]bool, numCols)
	}

	var dfs func(row, col, prev int, grid [][]bool)
	dfs = func(row, col, prev int, grid [][]bool) {
		if row < 0 || col < 0 || row == numRows || col == numCols {
			return
		}
		if grid[row][col] {
			return
		}
		if heights[row][col] < prev {
			return
		}

		grid[row][col] = true
		curr := heights[row][col]

		dfs(row+1, col, curr, grid)
		dfs(row-1, col, curr, grid)
		dfs(row, col+1, curr, grid)
		dfs(row, col-1, curr, grid)
	}

	// Loop through pacific
	for col := range numCols {
		dfs(0, col, 0, canReachFromPacific)
	}
	for row := range numRows {
		dfs(row, 0, 0, canReachFromPacific)
	}

	// Loop through atlantic
	for col := range numCols {
		dfs(numRows-1, col, 0, canReachFromAtlantic)
	}
	for row := range numRows {
		dfs(row, numCols-1, 0, canReachFromAtlantic)
	}

	result := make([][]int, 0)

	for row := range numRows {
		for col := range numCols {
			if canReachFromPacific[row][col] && canReachFromAtlantic[row][col] {
				coordinates := []int{row, col}
				result = append(result, coordinates)
			}
		}
	}

	return result
}
