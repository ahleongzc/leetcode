package backtracking

func subsets(nums []int) [][]int {
	output := make([][]int, 0)
	buffer := make([]int, 0)

	var dfs func(index int)
	dfs = func(index int) {
		bufferCopy := make([]int, len(buffer))
		copy(bufferCopy, buffer)
		output = append(output, bufferCopy)

		if index == len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			buffer = append(buffer, nums[i])
			dfs(i + 1)
			buffer = buffer[:len(buffer)-1]
		}
	}

	dfs(0)
	return output
}
