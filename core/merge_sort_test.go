package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	t.Run("1 element", func(t *testing.T) {
		nums := []int{1}
		res := MergeSort(nums)
		assert.IsIncreasing(t, res)
	})

	t.Run("5 elements", func(t *testing.T) {
		nums := []int{4, 3, 5, 1, 2}
		res := MergeSort(nums)
		assert.IsIncreasing(t, res)
	})
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		out := make([]int, len(nums))
		copy(out, nums)
		return out
	}

	mid := len(nums) / 2

	leftCopy := make([]int, mid)
	rightCopy := make([]int, len(nums)-mid)

	copy(leftCopy, nums[:mid])
	copy(rightCopy, nums[mid:])

	left := MergeSort(leftCopy)
	right := MergeSort(rightCopy)

	return conquer(left, right)
}

func conquer(a []int, b []int) []int {
	final := []int{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for i < len(a) {
		final = append(final, a[i])
		i++
	}

	for j < len(b) {
		final = append(final, b[j])
		j++
	}

	return final
}
