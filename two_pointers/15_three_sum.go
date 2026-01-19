package twopointers

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	output := make([][]int, 0)
	sort.Ints(nums)

	for curr := range len(nums) {
		if curr > 0 && nums[curr] == nums[curr-1] {
			continue
		}

		left, right := curr+1, len(nums)-1
		for left < right {
			currSum := nums[curr] + nums[left] + nums[right]
			if currSum > 0 {
				right--
			} else if currSum < 0 {
				left++
			} else {
				output = append(output, []int{nums[curr], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return output
}
