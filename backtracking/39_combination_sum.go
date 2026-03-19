package backtracking

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	buffer := make([]int, 0)

	var dfs func(index, total int)
	dfs = func(index, total int) {
		if index == len(candidates) || total > target {
			return
		}

		if total == target {
			bufferCopy := make([]int, len(buffer))
			copy(bufferCopy, buffer)
			res = append(res, bufferCopy)
			return
		}

		for i := index; i < len(candidates); i++ {
			buffer = append(buffer, candidates[i])
			dfs(i, total+candidates[i])
			buffer = buffer[:len(buffer)-1]
		}
	}

	dfs(0, 0)
	return res
}
