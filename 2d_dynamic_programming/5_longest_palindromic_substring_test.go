package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{
		{
			name:     "example 0",
			s:        "a",
			expected: []string{"a"},
		},
		{
			name:     "example 1",
			s:        "babad",
			expected: []string{"bab", "aba"},
		},
		{
			name:     "example 2",
			s:        "cbbd",
			expected: []string{"bb"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := longestPalindrome(test.s)
			if !assert.Contains(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
