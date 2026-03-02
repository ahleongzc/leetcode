package core

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	qs(nums, 0, len(nums)-1)
	return nums
}

// 4, 6, 3, 5

func qs(nums []int, start, end int) {
	if start >= end {
		return
	}

	pivot := nums[end]
	left := start

	for i := left; i < end; i++ {
		if nums[i] < pivot {
			tmp := nums[left]
			nums[left] = nums[i]
			nums[i] = tmp
			left++
		}
	}

	nums[end] = nums[left]
	nums[left] = pivot

	qs(nums, start, left-1)
	qs(nums, left+1, end)
}
