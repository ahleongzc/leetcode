package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseRobberII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{3, 4, 3},
			expected: 4,
		},
		{
			name:     "example 2",
			nums:     []int{2, 9, 8, 3, 6},
			expected: 15,
		},
		{
			name:     "example 3",
			nums:     []int{0},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := robII(test.nums)
			if !assert.Equal(t, output, test.expected) {
				t.Fatal()
			}
		})
	}
}
