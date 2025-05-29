package slidingwindow

func ContainNearbyDuplicates(nums []int, k int) bool {
	countMap := make(map[int]bool)
	i := 0

	for j, num := range nums {
		if j - i > k {
			delete(countMap, nums[i])
			i++
		}	

		if _, ok := countMap[num]; ok {
			return true
		}

		countMap[num] = true
	}

	return false
}
