package slidingwindow

func MinSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		}
		return 0
	}

	left := 0
	minLen := len(nums) + 1
	curr := 0

	for right := range len(nums) {
		curr += nums[right]
		for curr >= target {
			minLen = min(minLen, right-left+1)
			curr -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
