package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "example 1",
			piles:    []int{1, 4, 3, 2},
			h:        9,
			expected: 2,
		},
		{
			name:     "example 2",
			piles:    []int{25, 10, 23, 4},
			h:        4,
			expected: 25,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := minEatingSpeed(test.piles, test.h)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
