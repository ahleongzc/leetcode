package heappriorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		k, expected int
	}{
		{
			name:     "example 1",
			nums:     []int{2, 3, 1, 5, 4},
			k:        2,
			expected: 4,
		},
		{
			name:     "example 1",
			nums:     []int{2, 3, 1, 5, 4},
			k:        2,
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findKthLargest(test.nums, test.k)

			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
