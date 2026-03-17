package oneddp

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	for _, word := range wordDict {
		dict[word] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i > -1; i-- {
		for j := i; j < len(s); j++ {
			currWord := s[i : j+1]
			if _, ok := dict[currWord]; ok && dp[j+1] {
				dp[i] = true
			}
		}
	}

	return dp[0]
}
