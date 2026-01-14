package backtracking

// elements = 1, 2, 3, 4
// result = [1, 2, 3], [1, 2, 4], [1, 3, 4], [2, 3, 4]

func combine(n, k int) [][]int {
	output := make([][]int, 0)
	buffer := make([]int, 0)

	var dfs func(curr int)
	dfs = func(curr int) {
		if len(buffer) == k {
			newBuffer := make([]int, len(buffer))
			copy(newBuffer, buffer)
			output = append(output, newBuffer)
			return
		}

		if curr == n+1 {
			return
		}

		buffer = append(buffer, curr)
		dfs(curr + 1)

		buffer = buffer[:len(buffer)-1]
		dfs(curr + 1)
	}

	dfs(1)
	return output
}
