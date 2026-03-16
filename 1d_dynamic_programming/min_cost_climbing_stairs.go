package oneddp

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[len(cost)] = 0
	dp[len(cost)-1] = cost[len(cost)-1]
	dp[len(cost)-2] = cost[len(cost)-2]

	for i := len(cost) - 3; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+2], dp[i+1])
	}

	return min(dp[0], dp[1])
}
