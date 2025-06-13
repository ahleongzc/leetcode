package backtracking

import "slices"

func SubsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)

	subsets := make([][]int, 0)
	intermediary := make([]int, 0)

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			intermediaryCopy := make([]int, len(intermediary))
			copy(intermediaryCopy, intermediary)
			subsets = append(subsets, intermediaryCopy)
			return
		}

		intermediary = append(intermediary, nums[i])
		backtrack(i + 1)

		intermediary = intermediary[:len(intermediary)-1]

		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
		backtrack(i + 1)
	}

	backtrack(0)

	return subsets
}
