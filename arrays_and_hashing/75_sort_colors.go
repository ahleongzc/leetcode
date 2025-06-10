package arraysandhashing

func SortColors(nums []int) {
	merge := func(left, mid, right int) {
		leftSlice := make([]int, 0)
		rightSlice := make([]int, 0)

		for i := left; i <= mid; i++ {
			leftSlice = append(leftSlice, nums[i])
		}

		for i := mid + 1; i <= right; i++ {
			rightSlice = append(rightSlice, nums[i])
		}

		i, j := 0, 0

		for i < len(leftSlice) && j < len(rightSlice) {
			if leftSlice[i] <= rightSlice[j] {
				nums[left] = leftSlice[i]
				i++
			} else {
				nums[left] = rightSlice[j]
				j++
			}
			left++
		}

		for i < len(leftSlice) {
			nums[left] = leftSlice[i]
			i++
			left++
		}

		for j < len(rightSlice) {
			nums[left] = rightSlice[j]
			j++
			left++
		}
	}

	var mergeSort func(left, right int)
	mergeSort = func(left, right int) {
		if left >= right {
			return
		}

		mid := (right + left) / 2
		mergeSort(left, mid)
		mergeSort(mid+1, right)
		merge(left, mid, right)
	}

	mergeSort(0, len(nums)-1)
}
