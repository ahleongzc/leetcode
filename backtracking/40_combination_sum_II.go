package backtracking

import (
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	res := make([][]int, 0)
	buffer := make([]int, 0)

	var dfs func(index, total int)
	dfs = func(index, total int) {
		if total == target {
			bufferCopy := make([]int, len(buffer))
			copy(bufferCopy, buffer)
			res = append(res, bufferCopy)
			return
		}

		if index == len(candidates) || total > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			buffer = append(buffer, candidates[i])
			dfs(i+1, total+candidates[i])
			buffer = buffer[:len(buffer)-1]
		}
	}

	dfs(0, 0)
	return res
}
