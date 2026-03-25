package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name, s, p string
		expected   bool
	}{
		{
			name:     "example 0",
			s:        "a",
			p:        "a*",
			expected: true,
		},
		{
			name:     "example 1",
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			name:     "example 2",
			s:        "aa",
			p:        "*",
			expected: true,
		},
		{
			name:     "example 3",
			s:        "cb",
			p:        "?a",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isMatch(test.s, test.p)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal(test.name)
			}
		})
	}
}
