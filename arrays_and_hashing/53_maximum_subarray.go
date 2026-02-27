package arraysandhashing

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maximumSum := 0
	currSum := 0

	for _, num := range nums {
		currSum = max(currSum, 0)
		currSum += num
		maximumSum = max(maximumSum, currSum)
	}

	return maximumSum
}
