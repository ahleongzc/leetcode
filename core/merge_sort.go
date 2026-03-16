package core

func mergeSort(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}

	mid := len(numbers) / 2
	left := mergeSort(numbers[:mid])
	right := mergeSort(numbers[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	for l < len(left) {
		res = append(res, left[l])
		l++
	}

	for r < len(right) {
		res = append(res, right[r])
		r++
	}

	return res
}
