package backtracking

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	keyMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	res := make([]string, 0)

	var backtrack func(i int)
	intermediary := ""

	backtrack = func(i int) {
		if i == len(digits) {
			res = append(res, intermediary)
			return
		}

		for _, c := range keyMap[string(digits[i])] {
			intermediary += c
			backtrack(i + 1)
			intermediary = intermediary[:len(intermediary)-1]
		}
	}

	backtrack(0)

	return res
}
