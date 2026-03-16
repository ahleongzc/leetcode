package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeWays(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1",
			s:        "12",
			expected: 2,
		},
		{
			name:     "example 2",
			s:        "226",
			expected: 3,
		},
		{
			name:     "example 3",
			s:        "06",
			expected: 0,
		},
		{
			name:     "example 4",
			s:        "99",
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := numDecodings(test.s)
			if !assert.Equal(t, output, test.expected) {
				t.Fatal()
			}
		})
	}
}
