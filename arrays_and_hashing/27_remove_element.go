package arraysandhashing

func removeElement(nums []int, val int) int {
	curr, last := 0, len(nums)-1
	output := 0

	for curr <= last {
		if nums[curr] != val {
			output++
			curr++
			continue
		}

		nums[curr] = nums[last]
		nums[last] = val
		last--
	}

	return output
}
