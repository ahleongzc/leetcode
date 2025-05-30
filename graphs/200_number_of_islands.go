package graphs

func NumIslandBFS(grid [][]byte) int {
	num := 0
	num_rows := len(grid)
	num_cols := len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var bfs = func(row, col int) {
		queue := make([][]int, 0)
		queue = append(queue, []int{row, col})
		grid[row][col] = byte('0')

		for len(queue) > 0 {
			element := queue[0]
			queue = queue[1:]

			for _, val := range directions {
				delta_x := val[0]
				delta_y := val[1]

				new_x := delta_x + element[0]
				new_y := delta_y + element[1]
				

				if new_x < 0 || new_x == num_rows || new_y < 0 || new_y == num_cols || grid[new_x][new_y] == byte('0') {
					continue
				}

				grid[new_x][new_y] = byte('0')
				queue = append(queue, []int{new_x, new_y})
			}
		}
	}

	for row := range num_rows {
		for col := range num_cols {
			if grid[row][col] == byte('1') {
				bfs(row, col)
				num++
			}
		}
	}

	return num
}

func NumIslandsDFS(grid [][]byte) int {
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
				num++
			}
		}
	}

	return num
}
