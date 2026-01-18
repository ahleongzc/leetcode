package core

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	midPoint := len(numbers) / 2
	leftHalf := mergeSort(numbers[:midPoint])
	rightHalf := mergeSort(numbers[midPoint:])

	return merge(leftHalf, rightHalf)
}

func merge(leftHalf, rightHalf []int) []int {
	output := make([]int, 0)
	leftPtr, rightPtr := 0, 0

	for leftPtr < len(leftHalf) && rightPtr < len(rightHalf) {
		if leftHalf[leftPtr] <= rightHalf[rightPtr] {
			output = append(output, leftHalf[leftPtr])
			leftPtr++
		} else {
			output = append(output, rightHalf[rightPtr])
			rightPtr++
		}
	}

	for leftPtr < len(leftHalf) {
		output = append(output, leftHalf[leftPtr])
		leftPtr++
	}

	for rightPtr < len(rightHalf) {
		output = append(output, rightHalf[rightPtr])
		rightPtr++
	}

	return output
}
