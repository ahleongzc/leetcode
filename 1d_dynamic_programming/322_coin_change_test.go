package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "example 1",
			coins:    []int{1, 5, 10},
			amount:   12,
			expected: 3,
		},
		{
			name:     "example 2",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "example 3",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := coinChange(test.coins, test.amount)
			if !assert.Equal(t, test.expected, output) {
				t.Fatalf("test case %s", test.name)
			}
		})
	}
}
