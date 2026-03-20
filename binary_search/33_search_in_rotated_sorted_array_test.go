package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "example 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "example 3",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := search(test.nums, test.target)
			if !assert.Equal(t, actual, test.expected) {
				t.Fatalf("test case %v", test.name)
			}
		})
	}
}
