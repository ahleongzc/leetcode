package oneddp

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	words := make(map[string]struct{})
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			curr := s[i : j+1]
			// if the word exists in the dictionary and the subproblem is true
			// set the current entry to true AND break
			// or else the value will be overwritten by a longer string that exists but the subproblem is false
			if _, exists := words[curr]; exists && dp[j+1] {
				dp[i] = true
				break
			}
		}
	}

	return dp[0]
}
