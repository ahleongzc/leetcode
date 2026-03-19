package twoddp

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for amt := 0; amt < amount+1; amt++ {
			if amt < coin {
				continue
			}
			dp[amt] += dp[amt-coin]
		}
	}

	return dp[amount]
}
