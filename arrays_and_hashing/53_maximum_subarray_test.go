package arraysandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name: "example 1",
			nums: []int{
				-2, 1, -3, 4, -1, 2, 1, -5, 4,
			},
			expected: 6,
		},
		{
			name: "example 2",
			nums: []int{
				1,
			},
			expected: 1,
		},
		{
			name: "example 3",
			nums: []int{
				5, 4, -1, 7, 8,
			},
			expected: 23,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := maxSubArray(test.nums)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf(
					"test: %v - expected: %v, got %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
