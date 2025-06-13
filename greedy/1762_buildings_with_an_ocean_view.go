package greedy

import "slices"

func FindBuildings(heights []int) []int {
	result := make([]int, 0)
	result = append(result, len(heights)-1)

	for i := len(heights) - 2; i >= 0; i-- {
		if heights[i] > heights[i+1] {
			result = append(result, i)
		}
		heights[i] = max(heights[i], heights[i+1])
	}

	slices.Reverse(result)

	return result
}
