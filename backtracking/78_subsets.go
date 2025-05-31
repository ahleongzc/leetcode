package backtracking

func Subsets(nums []int) [][]int {
	res := make([][]int, 0)

	var dfs func(i int, intermediary []int)

	dfs = func(i int, intermediary []int) {
		if i == len(nums) {
			intermediaryCopy := make([]int, len(intermediary))
			copy(intermediaryCopy, intermediary)
			res = append(res, intermediaryCopy)
			return
		}

		intermediary = append(intermediary, nums[i])
		dfs(i+1, intermediary)

		intermediary = intermediary[:len(intermediary)-1]
		dfs(i+1, intermediary)
	}

	dfs(0, make([]int, 0))

	return res
}
