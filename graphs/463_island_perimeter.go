package graphs

func IslandPerimeterDFS(grid [][]int) int {
	perimeter := 0
	num_rows := len(grid)
	num_cols := len(grid[0])
	visited := make(map[[2]int]bool)

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if _, ok := visited[[2]int{row, col}]; ok {
			return 0
		}

		if row < 0 || col < 0 || row == num_rows || col == num_cols || grid[row][col] == 0 {
			return 1
		}

		grid[row][col] = 0
		visited[[2]int{row, col}] = true

		sum := 0
		sum += dfs(row + 1, col)
		sum += dfs(row - 1, col)
		sum += dfs(row, col + 1)
		sum += dfs(row, col - 1)

		return sum
	}

	for row := range num_rows {
		for col := range num_cols {
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
	num_rows := len(grid)
	num_cols := len(grid[0])

	var getPerimeter = func(row, col int) int {
		p := 0

		for _, direction := range directions {
			new_row := row + direction[0]
			new_col := col + direction[1]

			if new_row < 0 || new_row == num_rows || new_col < 0 || new_col == num_cols || grid[new_row][new_col] == 0 {
				p++
			}
		}

		return p
	}

	for row := range num_rows {
		for col := range num_cols {
			if grid[row][col] == 1 {
				perimeter += getPerimeter(row, col)
			}
		}
	}

	return perimeter
}
