package arraysandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, -2, 3, -2},
			expected: 3,
		},
		{
			name:     "example 2",
			nums:     []int{5, -3, 5},
			expected: 10,
		},
		{
			name:     "example 3",
			nums:     []int{-3, -2, -3},
			expected: -2,
		},
		{
			name:     "example 4",
			nums:     []int{1, 1, 1},
			expected: 3,
		},
		{
			name:     "example 5",
			nums:     []int{3, -1, 2, -1},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := maxSubarraySumCircular(test.nums)

			if !assert.Equal(t, output, test.expected) {
				t.Fatalf("test %v: expected %v, got %v", test.name, test.expected, output)
			}
		})
	}
}
