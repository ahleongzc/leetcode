package twoddp

func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := range len(s) {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i > -1; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}
