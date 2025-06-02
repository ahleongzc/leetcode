package twopointers

func Merge(nums1 []int, m int, nums2 []int, n int) {
	// TIP: Start from the back to prevent overriding the one at the front
	// A = [4, 5, 6, 0, 0, 0]
	// B = [1, 2, 3]

	endNumsOne := m - 1
	endNumsTwo := n - 1
	endIndex := m + n - 1

	for endNumsOne >= 0 && endNumsTwo >= 0 {
		if nums1[endNumsOne] > nums2[endNumsTwo] {
			nums1[endIndex] = nums1[endNumsOne]
			endNumsOne--
		} else {
			nums1[endIndex] = nums2[endNumsTwo]
			endNumsTwo--
		}
		endIndex--
	}

	for endNumsOne >= 0 {
		nums1[endIndex] = nums1[endNumsOne]
		endNumsOne--
		endIndex--
	}

	for endNumsTwo >= 0 {
		nums1[endIndex] = nums2[endNumsTwo]
		endNumsTwo--
		endIndex--
	}
}
