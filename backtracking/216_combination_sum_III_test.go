package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name     string
		k, n     int
		expected [][]int
	}{
		{
			name: "example 0",
			k:    9,
			n:    45,
			expected: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "example 1",
			k:    3,
			n:    7,
			expected: [][]int{
				{1, 2, 4},
			},
		},
		{
			name: "example 2",
			k:    3,
			n:    7,
			expected: [][]int{
				{1, 2, 4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := combinationSum3(test.k, test.n)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
