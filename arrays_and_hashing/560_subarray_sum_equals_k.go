package arraysandhashing

func subarraySum(nums []int, k int) int {
	output := 0
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	sum := 0

	for _, num := range nums {
		sum += num
		diff := sum - k

		if count, exists := prefixMap[diff]; exists {
			output += count
		}
		prefixMap[sum] += 1
	}

	return output
}
