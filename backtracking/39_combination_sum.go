package backtracking

func CombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	intermediary := make([]int, 0)

	var dfs func(i, total int)
	dfs = func(i, total int) {
		if i == len(candidates) || total > target {
			return
		}
		if total == target {
			intermediaryCopy := make([]int, len(intermediary))
			copy(intermediaryCopy, intermediary)
			res = append(res, intermediaryCopy)
			return
		}

		intermediary = append(intermediary, candidates[i])
		dfs(i, total+candidates[i])

		intermediary = intermediary[:len(intermediary)-1]
		dfs(i+1, total)
	}

	dfs(0, 0)

	return res
}
