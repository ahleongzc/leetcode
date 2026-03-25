package twoddp

func findTargetSumWays(nums []int, target int) int {
	dp := make(map[[2]int]int)

	var dfs func(i, currSum int) int
	dfs = func(i, currSum int) int {
		if i == len(nums) {
			if currSum == target {
				return 1
			}
			return 0
		}

		key := [2]int{i, currSum}
		if val, ok := dp[key]; ok {
			return val
		}

		left := dfs(i+1, currSum+nums[i])
		right := dfs(i+1, currSum-nums[i])

		dp[key] = left + right

		return left + right
	}

	return dfs(0, 0)
}
