package graphs

func IslandPerimeterDFS(grid [][]int) int {
	perimeter := 0
	numRows := len(grid)
	numCols := len(grid[0])
	visited := make(map[[2]int]bool)

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if _, ok := visited[[2]int{row, col}]; ok {
			return 0
		}

		if row < 0 || col < 0 || row == numRows || col == numCols || grid[row][col] == 0 {
			return 1
		}

		visited[[2]int{row, col}] = true

		sum := 0
		sum += dfs(row+1, col)
		sum += dfs(row-1, col)
		sum += dfs(row, col+1)
		sum += dfs(row, col-1)

		return sum
	}

	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == 1 {
				perimeter += dfs(row, col)
			}
		}
	}

	return perimeter
}

func IslandPerimeterIteration(grid [][]int) int {
	perimeter := 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	numRows := len(grid)
	numCols := len(grid[0])

	var getPerimeter = func(row, col int) int {
		p := 0

		for _, direction := range directions {
			newRow := row + direction[0]
			newCol := col + direction[1]

			if newRow < 0 || newRow == numRows || newCol < 0 || newCol == numCols || grid[newRow][newCol] == 0 {
				p++
			}
		}

		return p
	}

	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == 1 {
				perimeter += getPerimeter(row, col)
			}
		}
	}

	return perimeter
}
