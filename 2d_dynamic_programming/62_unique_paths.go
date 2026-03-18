package twoddp

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][n-1] = 1
	}

	for i := range n {
		dp[m-1][i] = 1
	}

	for row := m - 2; row > -1; row-- {
		for col := n - 2; col > -1; col-- {
			dp[row][col] = dp[row+1][col] + dp[row][col+1]
		}
	}

	return dp[0][0]
}
