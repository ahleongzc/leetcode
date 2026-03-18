package twoddp

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for row := range dp {
		dp[row] = make([]int, len(text1)+1)
	}

	for i := len(text2) - 1; i > -1; i-- {
		for j := len(text1) - 1; j > -1; j-- {
			if text1[j] == text2[i] {
				dp[i][j] = 1 + dp[i+1][j+1]
				continue
			}

			dp[i][j] = max(dp[i+1][j], dp[i][j+1])
		}
	}

	return dp[0][0]
}
