package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 3, 3},
			expected: 4,
		},
		{
			name:     "example 2",
			nums:     []int{2, 9, 8, 3, 6},
			expected: 16,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := rob(test.nums)
			if !assert.Equal(t, output, test.expected) {
				t.Fatal()
			}
		})
	}
}
