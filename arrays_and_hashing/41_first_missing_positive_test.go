package arraysandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name: "example 1",
			nums: []int{
				1, 2, 0,
			},
			expected: 3,
		},
		{
			name: "example 2",
			nums: []int{
				3, 4, -1, 1,
			},
			expected: 2,
		},
		{
			name: "example 3",
			nums: []int{
				7, 8, 9, 11, 12,
			},
			expected: 1,
		},
		{
			name: "example 4",
			nums: []int{
				100000, 3, 4000, 2, 15, 1, 99999,
			},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := firstMissingPositive(test.nums)
			if !assert.Equal(t, test.expected, output) {
				t.Fatalf(
					"test case %v - expected: %v, got: %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
