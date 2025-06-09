package arraysandhashing

func SortArray(nums []int) []int {
	merge := func(nums []int, left, mid, right int) {
		leftSlice := make([]int, 0)
		rightSlice := make([]int, 0)

		for i := left; i <= mid; i++ {
			leftSlice = append(leftSlice, nums[i])
		}

		for i := mid + 1; i <= right; i++ {
			rightSlice = append(rightSlice, nums[i])
		}

		ptrA := 0
		ptrB := 0

		for ptrA < len(leftSlice) && ptrB < len(rightSlice) {
			if leftSlice[ptrA] <= rightSlice[ptrB] {
				nums[left] = leftSlice[ptrA]
				ptrA++
			} else {
				nums[left] = rightSlice[ptrB]
				ptrB++
			}
			left++
		}

		for ptrA < len(leftSlice) {
			nums[left] = leftSlice[ptrA]
			left++
			ptrA++
		}

		for ptrB < len(rightSlice) {
			nums[left] = rightSlice[ptrB]
			left++
			ptrB++
		}
	}

	var mergeSort func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}

		mid := (left + right) / 2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}

	mergeSort(nums, 0, len(nums)-1)

	return nums
}
