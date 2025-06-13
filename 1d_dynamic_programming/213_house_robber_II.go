package oneddp

func RobII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(nums []int) int {
		dp := make([]int, len(nums))
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])

		for i := 2; i < len(nums); i++ {
			dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		}

		return dp[len(nums)-1]
	}

	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
