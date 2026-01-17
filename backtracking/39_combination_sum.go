package backtracking

func CombinationSum(candidates []int, target int) [][]int {
	output := make([][]int, 0)
	buffer := make([]int, 0)

	var dfs func(index, currSum int)
	dfs = func(index, currSum int) {
		if currSum > target {
			return
		}

		if currSum == target {
			newSlice := make([]int, len(buffer))
			copy(newSlice, buffer)
			output = append(output, newSlice)
			return
		}

		for i := index; i < len(candidates); i++ {
			buffer = append(buffer, candidates[i])
			dfs(i, currSum+candidates[i])
			buffer = buffer[:len(buffer)-1]
		}
	}

	dfs(0, 0)
	return output
}
