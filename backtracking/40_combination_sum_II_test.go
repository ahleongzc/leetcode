package backtracking

import (
	"reflect"
	"testing"
)

func TestCombinationSumII(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "example 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name:       "example 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected: [][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := combinationSum2(test.candidates, test.target)
			normalizeCombinations(output)
			normalizeCombinations(test.expected)

			if !reflect.DeepEqual(output, test.expected) {
				t.Fatalf(
					"test case: %v = expected: %v, got %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
