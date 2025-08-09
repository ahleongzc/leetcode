package arraysandhashing

func IsValidSudoku(board [][]byte) bool {
	hashSet := make(map[byte]bool)

	// check rows
	for _, row := range board {
		hashSet = make(map[byte]bool)
		for _, element := range row {
			if element == '.' {
				continue
			}
			if _, ok := hashSet[element]; ok {
				return false
			}
			hashSet[element] = true
		}
	}

	// check cols
	for col := 0; col < len(board[0]); col++ {
		hashSet = make(map[byte]bool)
		for row := range board {
			if board[row][col] == '.' {
				continue
			}
			if _, ok := hashSet[board[row][col]]; ok {
				return false
			}
			hashSet[board[row][col]] = true
		}
	}

	for row := range 3 {
		for col := range 3 {
			startRow := row * 3
			startCol := col * 3
			hashSet = make(map[byte]bool)

			for i := range 3 {
				for j := range 3 {
					currRow := startRow + i
					currCol := startCol + j
					if board[currRow][currCol] == '.' {
						continue
					}

					if _, ok := hashSet[board[currRow][currCol]]; ok {
						return false
					}
					hashSet[board[currRow][currCol]] = true
				}
			}
		}
	}

	return true
}
