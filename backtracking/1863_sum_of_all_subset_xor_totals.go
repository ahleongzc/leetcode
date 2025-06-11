package backtracking

func SubsetXORSum(nums []int) int {
	total := 0
	intermediary := make([]int, 0)

	totalXORFromSlice := func(slice []int) int {
		if len(slice) == 0 {
			return 0
		}
		total := slice[0]
		for i := 1; i < len(slice); i++ {
			total ^= slice[i]
		}
		return total
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			total += totalXORFromSlice(intermediary)
			return
		}

		intermediary = append(intermediary, nums[i])
		dfs(i + 1)

		intermediary = intermediary[:len(intermediary)-1]
		dfs(i + 1)
	}

	dfs(0)

	return total
}
