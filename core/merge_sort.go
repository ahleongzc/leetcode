package core

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	mid := len(numbers) / 2
	leftHalf := mergeSort(numbers[:mid])
	rightHalf := mergeSort(numbers[mid:])

	return merge(leftHalf, rightHalf)
}

func merge(left, right []int) []int {
	leftPtr, rightPtr := 0, 0
	result := make([]int, 0)

	for leftPtr < len(left) && rightPtr < len(right) {
		if left[leftPtr] < right[rightPtr] {
			result = append(result, left[leftPtr])
			leftPtr++
			continue
		}

		result = append(result, right[rightPtr])
		rightPtr++
	}

	for leftPtr < len(left) {
		result = append(result, left[leftPtr])
		leftPtr++
	}

	for rightPtr < len(right) {
		result = append(result, right[rightPtr])
		rightPtr++
	}

	return result
}
