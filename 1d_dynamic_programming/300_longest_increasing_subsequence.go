package oneddp

// Patience (Solitaire) sorting
func LengthOfLISOptimised(nums []int) int {
	patience := make([]int, 0)

	for _, num := range nums {
		left, right := 0, len(patience)

		for left < right {
			mid := (right + left) / 2
			if patience[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(patience) {
			patience = append(patience, num)
		} else {
			patience[left] = num
		}
	}

	return len(patience)
}

func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range len(nums) {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		curr := nums[i]

		for j := range len(nums[:i]) {
			prev := nums[j]

			if curr > prev {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	lis := 0
	for _, num := range dp {
		lis = max(lis, num)
	}

	return lis
}
