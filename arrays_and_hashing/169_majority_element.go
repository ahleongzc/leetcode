package arraysandhashing

func MajorityElement(nums []int) int {
	count := make(map[int]int, 0)

	for _, num := range nums {
		count[num]++
		if count[num] > len(nums)/2 {
			return num
		}
	}

	return nums[0]
}
