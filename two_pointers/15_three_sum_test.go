package twopointers

import (
	"leetcode/helper"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2}, {-1, 0, 1},
			},
		},
		{
			name:     "example 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name: "example 3",
			nums: []int{0, 0, 0},
			expected: [][]int{
				{0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := threeSum(test.nums)
			helper.Normalize2DIntSlices(output)
			helper.Normalize2DIntSlices(test.expected)

			if !reflect.DeepEqual(output, test.expected) {
				t.Fatalf(
					"test case: %v - expected: %v, got %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
