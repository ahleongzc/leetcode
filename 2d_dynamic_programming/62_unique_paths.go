package twoddp

func UniquePaths(m int, n int) int {
	board := make([][]int, m)
	for i := range board {
		board[i] = make([]int, n)
	}

	for row := range m {
		board[row][n-1] = 1
	}

	for col := range n {
		board[m-1][col] = 1
	}

	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			board[row][col] = board[row+1][col] + board[row][col+1]
		}
	}

	return board[0][0]
}
