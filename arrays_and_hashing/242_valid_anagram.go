package arraysandhashing

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runeCount := make(map[rune]int)

	for _, r := range s {
		runeCount[r]++
	}

	for _, r := range t {
		runeCount[r]--

		if runeCount[r] < 0 {
			return false
		}
	}

	return true
}