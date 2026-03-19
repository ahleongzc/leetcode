package oneddp

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	
	// this is permutation
	// 2, 1 -> 3
	// 1, 2 -> 3
	for i := 1; i < target+1; i++ {
		for _, num := range nums {
			if num > i {
				continue
			}
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
}

func combinationSum4Recursion(nums []int, target int) int {
	totalCombinations := 0

	var dfs func(currSum int)
	dfs = func(currSum int) {
		if currSum == target {
			totalCombinations++
			return
		}

		if currSum > target {
			return
		}

		for _, num := range nums {
			dfs(currSum + num)
		}
	}

	dfs(0)
	return totalCombinations
}
