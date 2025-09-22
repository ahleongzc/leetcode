package graphs

// WallsAndGates fills each empty room with the distance to its nearest gate.
// -1 represents a wall
// 0 represents a gate
// INF (2147483647) represents an empty room

const INF = 2147483647

func WallsAndGates(rooms [][]int) {
	ROWS := len(rooms)
	COLS := len(rooms[0])

	queue := make([][2]int, 0)
	directions := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	// populate the queue
	for row := range len(rooms) {
		for col := range len(rooms[0]) {
			if rooms[row][col] == 0 {
				queue = append(queue, [2]int{row, col})
			}
		}
	}

	// bfs
	for len(queue) != 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, direction := range directions {
			dx, dy := direction[0], direction[1]
			newRow, newCol := row+dx, col+dy

			// out of bounds
			if newRow == ROWS || newCol == COLS || newRow < 0 || newCol < 0 {
				continue
			}

			// encounter a wall
			if rooms[newRow][newCol] == -1 {
				continue
			}
			
			// add to queue if it's an empty room
			if rooms[newRow][newCol] == INF {
				queue = append(queue, [2]int{newRow, newCol})
			}

			rooms[newRow][newCol] = min(rooms[row][col]+1, rooms[newRow][newCol])
		}
	}
}
