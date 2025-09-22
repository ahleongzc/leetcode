package binarysearch

func FindMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] < nums[right] {
			return nums[left]
		}

		mid := (right + left) / 2
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
