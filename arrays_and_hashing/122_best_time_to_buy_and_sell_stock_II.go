package arraysandhashing

func MaxProfitO1Space(prices []int) int {
	total := 0

	for i := 1; i < len(prices); i++ {
		dayToDayProfit := prices[i] - prices[i-1]
		if dayToDayProfit > 0 {
			total += dayToDayProfit
		}
	}

	return total
}

func MaxProfit(prices []int) int {
	total := 0
	dayToDayProfit := make([]int, len(prices))
	dayToDayProfit[0] = 0

	for i := 1; i < len(prices); i++ {
		dayToDayProfit[i] = prices[i] - prices[i-1]
	}

	for _, profit := range dayToDayProfit {
		if profit > 0 {
			total += profit
		}
	}

	return total
}
