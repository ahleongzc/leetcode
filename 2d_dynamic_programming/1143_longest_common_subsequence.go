package twoddp

func LongestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for row := range dp {
		dp[row] = make([]int, len(text2)+1)
	}

	for row := len(text1) - 1; row >= 0; row-- {
		for col := len(text2) - 1; col >= 0; col-- {
			if text1[row] == text2[col] {
				dp[row][col] = dp[row+1][col+1] + 1
			} else {
				dp[row][col] = max(dp[row][col+1], dp[row+1][col])
			}
		}
	}

	return dp[0][0]
}
