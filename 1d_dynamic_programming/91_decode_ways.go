package oneddp

func numDecodings(s string) int {
	count := make(map[int]int)
	count[len(s)] = 1

	for i := len(s) - 1; i > -1; i-- {
		if s[i] == '0' {
			count[i] = 0
		} else {
			count[i] = count[i+1]
		}

		if i < len(s)-1 && (s[i] == '1' || (s[i] == '2' && s[i+1] < '7')) {
			count[i] += count[i+2]
		}
	}

	return count[0]
}

func numDecodingsRecursion(s string) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		res := 0
		if i == len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}

		res += dfs(i + 1)
		if i < len(s)-1 {
			if s[i] == '1' || (s[i] == '2' && s[i+1] < '7') {
				res += dfs(i + 2)
			}
		}

		return res
	}

	return dfs(0)
}
