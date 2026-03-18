package core

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	quickSortHelper(nums, 0, len(nums)-1)
	return nums
}

func quickSortHelper(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := nums[end]
	left := start

	for i := start; i < end; i++ {
		if nums[i] >= pivot {
			continue
		}
		nums[i], nums[left] = nums[left], nums[i]
		left++
	}

	nums[end] = nums[left]
	nums[left] = pivot

	quickSortHelper(nums, start, left-1)
	quickSortHelper(nums, left+1, end)
}
