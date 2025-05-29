package arraysandhashing

func LongestCommonPrefix(strs []string) string {
	lcp := ""
	referenceString := strs[0]

	for i, c := range referenceString {
		for _, str := range strs {
			if len(str) == i {
				return lcp
			}
			if str[i] != referenceString[i] {
				return lcp
			}
		}

		lcp += string(c)
	}

	return lcp
}