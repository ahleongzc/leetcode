package twopointers

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left := 0
	length := 0
	set := make(map[byte]bool)

	for right := range len(s) {
		for set[s[right]] {
			delete(set, s[left])
			left++
		}

		length = max(length, right-left+1)
		set[s[right]] = true
	}

	return length
}
