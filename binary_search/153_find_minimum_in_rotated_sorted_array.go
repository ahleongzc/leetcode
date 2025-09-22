package binarysearch

func FindMin(nums []int) int {
	res := nums[0]
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] < nums[right] {
			res = min(res, nums[left])
			return res
		}

		mid := (right + left) / 2
		res = min(res, nums[mid])
		
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
