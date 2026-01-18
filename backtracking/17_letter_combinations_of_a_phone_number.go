package backtracking

func letterCombinations(digits string) []string {
	combinations := make(map[byte][]string)
	combinations['2'] = []string{"a", "b", "c"}
	combinations['3'] = []string{"d", "e", "f"}
	combinations['4'] = []string{"g", "h", "i"}
	combinations['5'] = []string{"j", "k", "l"}
	combinations['6'] = []string{"m", "n", "o"}
	combinations['7'] = []string{"p", "q", "r", "s"}
	combinations['8'] = []string{"t", "u", "v"}
	combinations['9'] = []string{"w", "x", "y", "z"}

	output := make([]string, 0)

	var dfs func(digitIndex int, curr string)
	dfs = func(digitIndex int, curr string) {
		if digitIndex == len(digits) {
			output = append(output, curr)
			return
		}

		for _, char := range combinations[digits[digitIndex]] {
			curr += char
			dfs(digitIndex+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, "")
	return output
}
