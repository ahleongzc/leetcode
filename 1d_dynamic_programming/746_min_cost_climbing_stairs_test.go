package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "example 1",
			cost:     []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "example 2",
			cost:     []int{1, 2, 1, 2, 1, 1, 1},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := minCostClimbingStairs(test.cost)
			if !assert.Equal(t, output, test.expected) {
				t.Fatal()
			}
		})
	}
}
