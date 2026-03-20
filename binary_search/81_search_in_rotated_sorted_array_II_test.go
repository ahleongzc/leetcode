package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInRotatedSortedArray2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "example 1",
			nums:     []int{3, 4, 4, 5, 6, 1, 2, 2},
			target:   1,
			expected: true,
		},
		{
			name:     "example 2",
			nums:     []int{3, 5, 6, 0, 0, 1, 2},
			target:   4,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := search2(test.nums, test.target)
			if !assert.Equal(t, actual, test.expected) {
				t.Fatalf("test case %v", test.name)
			}
		})
	}
}
