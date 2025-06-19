package arraysandhashing

func LongestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	longest := 0
	for num := range numMap {
		// The start of a sequence doesn't have the num-1 inside the numMap
		if _, exists := numMap[num-1]; exists {
			continue
		}

		consecutive := 0
		for {
			if _, ok := numMap[num+consecutive]; !ok {
				break
			}
			consecutive++
		}
		longest = max(longest, consecutive)
	}

	return longest
}
