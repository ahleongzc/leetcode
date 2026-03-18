package twoddp

func longestPalindrome(s string) string {
	n := len(s)
	resIdx, resLen := 0, 0

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// the dp table represents the substring starting from i and ending at j
	// 		j →
	//         0      1       2        3         4
	// i ↓
	// 0     "b"    "ba"    "bab"    "baba"    "babab"
	// 1             "a"     "ab"     "aba"     "abab"
	// 2                      "b"     "ba"      "bab"
	// 3                               "a"      "ab"
	// 4                                        "b"

	for row := n - 1; row >= 0; row-- {
		for col := row; col < n; col++ {
			if s[row] == s[col] && (col-row <= 2 || dp[row+1][col-1]) {
				dp[row][col] = true
				if resLen < (col - row + 1) {
					resIdx = row
					resLen = col - row + 1
				}
			}
		}
	}

	return s[resIdx : resIdx+resLen]
}

func longestPalindromeTwoPointers(s string) string {
	res := ""

	var checkPalindromeFromCenter func(left, right int)
	checkPalindromeFromCenter = func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}

	for i := range s {
		// check odd length
		checkPalindromeFromCenter(i, i)
		// check even length
		checkPalindromeFromCenter(i, i+1)
	}

	return res
}
