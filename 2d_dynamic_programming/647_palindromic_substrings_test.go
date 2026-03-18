package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 0",
			s:        "a",
			expected: 1,
		},
		{
			name:     "example 1",
			s:        "abc",
			expected: 3,
		},
		{
			name:     "example 2",
			s:        "aaa",
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := countSubstrings(test.s)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
