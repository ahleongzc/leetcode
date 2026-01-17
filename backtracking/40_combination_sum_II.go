package backtracking

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	output := make([][]int, 0)
	buffer := make([]int, 0)

	candidatesCopy := make([]int, len(candidates))
	copy(candidatesCopy, candidates)
	slices.Sort(candidatesCopy)

	var dfs func(index, currSum int)
	dfs = func(index, currSum int) {
		if currSum > target {
			return
		}

		if currSum == target {
			bufferCopy := make([]int, len(buffer))
			copy(bufferCopy, buffer)
			output = append(output, bufferCopy)
			return
		}

		for i := index; i < len(candidatesCopy); i++ {
			// skips duplicate at the same recursion depth
			if i > index && candidatesCopy[i] == candidatesCopy[i-1] {
				continue
			}
			buffer = append(buffer, candidatesCopy[i])
			dfs(i+1, currSum+candidatesCopy[i])
			buffer = buffer[:len(buffer)-1]
		}
	}

	dfs(0, 0)
	return output
}
