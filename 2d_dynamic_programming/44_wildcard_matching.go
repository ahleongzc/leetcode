package twoddp

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[len(p)][len(s)] = true

	for row := len(p) - 1; row > -1; row-- {
		if p[row] == '*' {
			for col := len(s); col > -1; col-- {
				dp[row][col] = true
			}
		} else {
			break
		}
	}

	for row := len(p) - 1; row > -1; row-- {
		for col := len(s) - 1; col > -1; col-- {
			if p[row] == '?' || p[row] == s[col] {
				dp[row][col] = dp[row+1][col+1]
				continue
			}
			if p[row] == '*' {
				dp[row][col] = dp[row+1][col+1] || dp[row+1][col] || dp[row][col+1]
				continue
			}

			dp[row][col] = false
		}
	}

	return dp[0][0]
}
