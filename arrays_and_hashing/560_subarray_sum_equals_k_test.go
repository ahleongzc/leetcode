package arraysandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "example 2",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			name:     "example 3",
			nums:     []int{1},
			k:        0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := subarraySum(test.nums, test.k)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf("test case: %v - expected: %v, got: %v", test.name, test.expected, output)
			}
		})
	}
}
