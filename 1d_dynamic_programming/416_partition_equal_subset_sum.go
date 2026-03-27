package oneddp

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make(map[int]struct{})
	dp[0] = struct{}{}

	for _, num := range nums {
		next := make(map[int]struct{})
		for key := range dp {
			if key == target {
				return true
			}
			if key+num == target {
				return true
			}
			next[key] = struct{}{}
			next[key+num] = struct{}{}
		}
		dp = next
	}

	return false
}

func canPartitionMemoize(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make(map[[2]int]bool)

	var dfs func(i, curr int) bool
	dfs = func(i, curr int) bool {
		key := [2]int{i, curr}
		if val, ok := dp[key]; ok {
			return val
		}

		if curr > target {
			return false
		}

		if curr == target {
			return true
		}

		if i == len(nums) {
			return false
		}

		include := dfs(i+1, curr+nums[i])
		notInclude := dfs(i+1, curr)

		dp[key] = include || notInclude
		return dp[key]
	}

	return dfs(0, 0)
}
