package oneddp

func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	skipFirst := robLinear(nums[1:])
	skipLast := robLinear(nums[:len(nums)-1])

	return max(skipFirst, skipLast)
}

func robLinear(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[len(nums)] = 0
	dp[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 2; i > -1; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}

	return dp[0]
}
