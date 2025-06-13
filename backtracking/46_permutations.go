package backtracking

import (
	"slices"
)

func Permute(nums []int) [][]int {

	var backtrack func(i int) [][]int
	backtrack = func(i int) [][]int {
		if i == len(nums) {
			return [][]int{{}}
		}

		permutations := backtrack(i + 1)
		result := make([][]int, 0)

		for _, permutation := range permutations {
			for pos := range len(permutation) + 1 {
				newPermutation := make([]int, len(permutation))
				copy(newPermutation, permutation)

				newPermutation = slices.Insert(newPermutation, pos, nums[i])
				result = append(result, newPermutation)
			}
		}

		return result
	}

	return backtrack(0)
}
