package graphs

func NumIslandBFS(grid [][]byte) int {
	num := 0
	numRows := len(grid)
	numCols := len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var bfs = func(row, col int) {
		queue := make([][]int, 0)
		queue = append(queue, []int{row, col})
		grid[row][col] = byte('0')

		for len(queue) > 0 {
			element := queue[0]
			queue = queue[1:]

			for _, val := range directions {
				deltaX := val[0]
				deltaY := val[1]

				newX := deltaX + element[0]
				newY := deltaY + element[1]

				if newX < 0 || newX == numRows || newY < 0 || newY == numCols || grid[newX][newY] == byte('0') {
					continue
				}

				grid[newX][newY] = byte('0')
				queue = append(queue, []int{newX, newY})
			}
		}
	}

	for row := range numRows {
		for col := range numCols {
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
	numRows := len(grid)
	numCols := len(grid[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row == numRows || col < 0 || col == numCols || grid[row][col] == byte('0') {
			return
		}

		grid[row][col] = byte('0')
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == byte('1') {
				dfs(row, col)
				num++
			}
		}
	}

	return num
}
