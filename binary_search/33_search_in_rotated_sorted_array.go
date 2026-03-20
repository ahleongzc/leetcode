package binarysearch

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// [4, 5, 6, 7, 0, 1, 2]
	//  L        M        R

	for left <= right {
		mid := (right + left) / 2
		if target == nums[mid] {
			return mid
		}

		// mid belongs to left half which is sorted
		if nums[mid] >= nums[left] {
			if target < nums[left] || target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// mid belongs to right half which is sorted
			if target > nums[right] || target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
