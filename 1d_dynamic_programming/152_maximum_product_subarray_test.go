package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name: "example 1",
			nums: []int{
				2, 3, -2, 4,
			},
			expected: 6,
		},
		{
			name: "example 2",
			nums: []int{
				-2, 0, -1,
			},
			expected: 0,
		},
		{
			name: "example 3",
			nums: []int{
				-2, 3, -4,
			},
			expected: 24,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := maxProduct(test.nums)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf(
					"test case: %v - expected: %v, got %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
