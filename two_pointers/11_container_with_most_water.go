package twopointers

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0

	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]

		area = max(area, min(leftHeight, rightHeight)*(right-left))

		if leftHeight <= rightHeight {
			left++
		} else {
			right--
		}
	}

	return area
}
