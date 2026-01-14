package arraysandhashing

func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		currentSum := numbers[left] + numbers[right]
		if currentSum == target {
			return []int{left + 1, right + 1}
		}

		if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
