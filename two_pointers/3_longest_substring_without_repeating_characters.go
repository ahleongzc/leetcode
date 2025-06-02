package twopointers

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	// Step 1: Initialization
	length := 1
	left := 0
	right := 1

	uniqueSet := make(map[byte]int)
	uniqueSet[s[left]] = 0

	// Step 2: Two pointers
	for right < len(s) {
		prevIndex, exists := uniqueSet[s[right]]
		if exists && left <= prevIndex {
			left = prevIndex + 1
		}

		length = max(length, right-left+1)
		uniqueSet[s[right]] = right
		right++
	}

	return length
}
