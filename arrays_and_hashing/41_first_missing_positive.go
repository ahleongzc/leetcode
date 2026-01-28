package arraysandhashing

func firstMissingPositive(nums []int) int {
	// the possible range of answers is [1, n + 1]
	// 1st step: make all the negative numbers become 0 since negative numbers cannot be part of the answer
	// this cleans up the array for us to perform negative marking since now the array only contains positive values
	for i, num := range nums {
		if num < 0 {
			nums[i] = 0
		}
	}

	for _, num := range nums {
		// the mapping of index to value is: value = index - 1
		// e.g. value of 1 = 0 index
		val := abs(num)

		// this if statement checks whether the value is within the possible set of answers
		if 1 <= val && val <= len(nums) {
			index := val - 1

			if nums[index] == 0 {
				nums[index] = -(len(nums) + 1)
			} else if nums[index] > 0 {
				nums[index] *= -1
			}
		}
	}

	for i, num := range nums {
		if num >= 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}
