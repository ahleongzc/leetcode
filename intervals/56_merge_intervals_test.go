package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name: "example 1",
			intervals: [][]int{
				{1, 3}, {2, 6}, {8, 10}, {15, 18},
			},
			expected: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
		{
			name: "example 2",
			intervals: [][]int{
				{1, 4}, {4, 5},
			},
			expected: [][]int{
				{1, 5},
			},
		},
		{
			name: "example 3",
			intervals: [][]int{
				{4, 7}, {1, 4},
			},
			expected: [][]int{
				{1, 7},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := merge(test.intervals)
			if !assert.ElementsMatch(t, test.expected, actual) {
				t.Fatal(test.name)
			}
		})
	}
}
