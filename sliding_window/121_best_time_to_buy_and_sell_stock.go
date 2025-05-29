package slidingwindow

func MaxProfit(prices []int) int {
	profit := 0
	i := 0

	for j := range len(prices) {
		if prices[j]-prices[i] < 0 {
			i = j
		}

		if prices[j]-prices[i] > profit {
			profit = prices[j] - prices[i]
		}
	}

	return profit
}
