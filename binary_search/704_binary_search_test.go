package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name: "example 1",
			nums: []int{
				-1, 0, 2, 4, 6, 8,
			},
			target:   4,
			expected: 3,
		},
		{
			name: "example 2",
			nums: []int{
				-1, 0, 2, 4, 6, 8,
			},
			target:   3,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := bsearch(test.nums, test.target)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
