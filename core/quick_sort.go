package core

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	qs(nums, 0, len(nums)-1)
	return nums
}

func qs(nums []int, start, end int) {
	if start >= end {
		return
	}

	pivot := nums[end]
	left := start

	for i := left; i < end; i++ {
		if nums[i] >= pivot {
			continue
		}

		tmp := nums[i]
		nums[i] = nums[left]
		nums[left] = tmp
		left++
	}

	nums[end] = nums[left]
	nums[left] = pivot

	qs(nums, start, left-1)
	qs(nums, left+1, end)
}
