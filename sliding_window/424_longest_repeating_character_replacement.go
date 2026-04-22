package slidingwindow

import "slices"

func characterReplacement(s string, k int) int {
	charCount := make([]int, 26)
	longestSubstringLength := 0
	left := 0

	for right, char := range s {
		charCount[int(char)%26]++
		for (right-left+1)-slices.Max(charCount) > k {
			charCount[int(s[left])%26]--
			left++
		}
		longestSubstringLength = max(longestSubstringLength, right-left+1)
	}

	return longestSubstringLength
}
