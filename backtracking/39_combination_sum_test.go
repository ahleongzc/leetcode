package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected [][]int
	}{
		{
			name:   "example 1",
			input:  []int{2, 3, 6, 7},
			target: 7,
			expected: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			name:   "example 2",
			input:  []int{2, 3, 5},
			target: 8,
			expected: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name:     "example 3",
			input:    []int{2},
			target:   1,
			expected: [][]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := CombinationSum(test.input, test.target)
			normalizeCombinations(output)
			normalizeCombinations(test.expected)

			if !reflect.DeepEqual(test.expected, output) {
				t.Fatalf(
					"test case: %v - expected: %v but got %v instead",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}

func normalizeCombinations(combos [][]int) {
	for _, c := range combos {
		sort.Ints(c)
	}

	sort.Slice(combos, func(i, j int) bool {
		a, b := combos[i], combos[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
}
