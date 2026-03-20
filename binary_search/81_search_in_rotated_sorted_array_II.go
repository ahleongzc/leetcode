package binarysearch

func search2(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return true
		}

		// [4, 5, 6, 7, 0, 1, 2]
		//  L        M        R

		// mid is in left sorted portion of the array
		if nums[mid] > nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// the reason why you shift the left pointer is because it is a sorted array
			// [ 2, 2, 2, 3, 1, 2 ] will soon become [ 3, 1, 2]

			// if you decrease from the right, you will risk removing the element that you want to find
			// [ 2, 2, 2, 3, 1, 2] will soon become [ 2, 2, 2, 3 ]
			left++
		}
	}

	return false
}
