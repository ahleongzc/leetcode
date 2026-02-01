package oneddp

import (
	"slices"
)

func maxProduct(nums []int) int {
	output := slices.Max(nums)
	currMin, currMax := 1, 1

	for _, num := range nums {
		temp := currMax * num
		currMax = max(num*currMax, num*currMin, num)
		currMin = min(temp, num*currMin, num)
		output = max(output, currMax)
	}

	return output
}
