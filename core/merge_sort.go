package core

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	mid := len(numbers) / 2
	left := mergeSort(numbers[:mid])
	right := mergeSort(numbers[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	leftPtr, rightPtr := 0, 0

	for leftPtr < len(left) && rightPtr < len(right) {
		if left[leftPtr] < right[rightPtr] {
			res = append(res, left[leftPtr])
			leftPtr++
		}
		res = append(res, right[rightPtr])
		rightPtr++
	}

	for leftPtr < len(left) {
		res = append(res, left[leftPtr])
		leftPtr++
	}

	for rightPtr < len(right) {
		res = append(res, right[rightPtr])
		rightPtr++
	}

	return res
}
