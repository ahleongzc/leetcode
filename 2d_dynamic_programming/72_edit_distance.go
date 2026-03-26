package twoddp

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := range dp {
		dp[i] = make([]int, len(word1)+1)
		dp[i][len(word1)] = len(word2) - i
	}

	for i := range len(word1) + 1 {
		dp[len(word2)][i] = len(word1) - i
	}

	for row := len(word2) - 1; row > -1; row-- {
		for col := len(word1) - 1; col > -1; col-- {
			if word1[col] == word2[row] {
				dp[row][col] = dp[row+1][col+1]
				continue
			}

			dp[row][col] = 1 + min(
				dp[row+1][col],
				dp[row][col+1],
				dp[row+1][col+1],
			)
		}
	}

	return dp[0][0]
}
