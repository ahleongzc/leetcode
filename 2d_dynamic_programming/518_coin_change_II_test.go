package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		coins    []int
		expected int
	}{
		{
			name:     "example 1",
			amount:   5,
			coins:    []int{1, 2, 5},
			expected: 4,
		},
		{
			name:     "example 2",
			amount:   3,
			coins:    []int{2},
			expected: 0,
		},
		{
			name:     "example 2",
			amount:   10,
			coins:    []int{10},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := change(test.amount, test.coins)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
