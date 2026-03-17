package oneddp

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)
	dp = append(dp, nums[0])

	var binarySearch func(target, start, end int) int
	binarySearch = func(target, start, end int) int {
		for start < end {
			mid := (start + end) / 2
			if dp[mid] < target {
				start = mid + 1
			} else {
				end = mid
			}
		}
		return start
	}

	for _, num := range nums {
		if num > dp[len(dp)-1] {
			dp = append(dp, num)
			continue
		}

		// replace the smallest possible number in dp where the number >= num
		id := binarySearch(num, 0, len(dp)-1)
		dp[id] = num
	}

	return len(dp)
}

func lengthOfLISQuadratic(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := range len(nums) {
		for j := range i {
			if nums[i] <= nums[j] {
				continue
			}
			dp[i] = max(dp[i], 1+dp[j])
		}
	}

	output := -1
	for _, val := range dp {
		output = max(output, val)
	}

	return output
}
