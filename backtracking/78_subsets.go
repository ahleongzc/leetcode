package backtracking

func subsets(nums []int) [][]int {
	output := make([][]int, 0)
	buffer := make([]int, 0)

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			newBuffer := make([]int, len(buffer))
			copy(newBuffer, buffer)
			output = append(output, newBuffer)
			return
		}

		buffer = append(buffer, nums[index])
		index++
		dfs(index)

		buffer = buffer[:len(buffer)-1]
		dfs(index)
	}

	dfs(0)
	return output
}
