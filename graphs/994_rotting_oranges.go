package graphs

func OrangesRotting(grid [][]int) int {
	// Step 1: Set up
	numRows := len(grid)
	numCols := len(grid[0])
	queue := make([][2]int, 0)
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Step 2: Add the initial rotten oranges into the queue
	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == 2 {
				queue = append(queue, [2]int{row, col})
			}
		}
	}

	// Step 3: BFS
	time := 0

	for len(queue) != 0 {
		rottenOranges := queue[:]
		queue = [][2]int{}

		for _, rottenOrange := range rottenOranges {
			for _, direction := range directions {
				newRow := direction[0] + rottenOrange[0]
				newCol := direction[1] + rottenOrange[1]

				// If out of bounds, continue
				if newRow < 0 || newRow == numRows || newCol < 0 || newCol == numCols {
					continue
				}

				// If it is an empty cell or the orange has already rotten, continue
				if grid[newRow][newCol] == 0 || grid[newRow][newCol] == 2 {
					continue
				}

				queue = append(queue, [2]int{newRow, newCol})
				grid[newRow][newCol] = 2
			}
		}

		// Only increment where there are new oranges to rot
		if len(queue) > 0 {
			time++
		}
	}

	for row := range numRows {
		for col := range numCols {
			if grid[row][col] == 1 {
				return -1
			}
		}
	}

	return time
}
