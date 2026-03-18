package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name           string
		m, n, expected int
	}{
		{
			name:     "example 1",
			m:        3,
			n:        6,
			expected: 21,
		},
		{
			name:     "example 2",
			m:        3,
			n:        3,
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := uniquePaths(test.m, test.n)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
