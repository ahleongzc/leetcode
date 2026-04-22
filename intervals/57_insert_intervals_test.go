package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertIntervals(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name: "example 1",
			intervals: [][]int{
				{1, 3}, {6, 9},
			},
			newInterval: []int{2, 5},
			expected: [][]int{
				{1, 5}, {6, 9},
			},
		},
		{
			name: "example 2",
			intervals: [][]int{
				{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
			},
			newInterval: []int{4, 8},
			expected: [][]int{
				{1, 2}, {3, 10}, {12, 16},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := insert(test.intervals, test.newInterval)
			if !assert.ElementsMatch(t, test.expected, actual) {
				t.Fatal(test.name)
			}
		})
	}
}
