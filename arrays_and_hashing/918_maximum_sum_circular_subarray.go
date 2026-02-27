package arraysandhashing

func maxSubarraySumCircular(nums []int) int {
	currMax := 0
	globalMax := nums[0]

	currMin := 0
	globalMin := nums[0]

	total := 0

	for _, num := range nums {
		currMax = max(currMax, 0)
		currMax += num
		globalMax = max(globalMax, currMax)

		currMin += num
		currMin = min(currMin, num)
		globalMin = min(globalMin, currMin)

		total += num
	}

	if globalMax < 0 {
		return globalMax
	}

	return max(total-globalMin, globalMax)
}
