package slidingwindow

func ContainNearbyDuplicates(nums []int, k int) bool {
	start := 0
	end := k

	countMap := make(map[int]bool)

	// First build the map
	for i := range k + 1 {
		if i == len(nums) {
			return false
		}

		if _, ok := countMap[nums[i]]; ok {
			return true
		}
		countMap[nums[i]] = true
	}

	for end < len(nums) {
		delete(countMap, nums[start])
		start++
		end++

		if end == len(nums) {
			return false
		}

		if _, ok := countMap[nums[end]]; ok {
			return true
		}
		countMap[nums[end]] = true
	}

	return false
}
