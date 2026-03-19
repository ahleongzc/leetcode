package backtracking

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	buffer := make([]int, 0)

	var dfs func(curr, total int)
	dfs = func(curr, total int) {
		if total == n {
			if len(buffer) == k {
				bufferCopy := make([]int, len(buffer))
				copy(bufferCopy, buffer)
				res = append(res, bufferCopy)
			}
			return
		}

		if curr > 9 || total > n {
			return
		}

		for i := curr; i < 10; i++ {
			if i == 0 {
				continue
			}
			buffer = append(buffer, i)
			dfs(i+1, total+i)
			buffer = buffer[:len(buffer)-1]
		}
	}

	dfs(0, 0)
	return res
}
