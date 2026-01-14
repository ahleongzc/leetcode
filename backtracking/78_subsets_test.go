package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "example 1",
			input: []int{1, 2, 3},
			expected: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3}},
		},
		{
			name:  "example 2",
			input: []int{0},
			expected: [][]int{
				{},
				{0},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := subsets(test.input)
			if !assert.ElementsMatch(t, result, test.expected) {
				t.Fatalf("%v: expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}
