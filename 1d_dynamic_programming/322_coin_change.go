package oneddp

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			}

			dp[i] = min(1+dp[i-coin], dp[i])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
