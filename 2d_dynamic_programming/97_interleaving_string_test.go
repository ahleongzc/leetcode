package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name, s1, s2, s3 string
		expected         bool
	}{
		{
			name:     "ex 1",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbcbcac",
			expected: true,
		},
		{
			name:     "ex 2",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbbaccc",
			expected: false,
		},
		{
			name:     "ex 3",
			s1:       "",
			s2:       "",
			s3:       "",
			expected: true,
		},
	}

	for _, test := range tests {
		actual := isInterleave(test.s1, test.s2, test.s3)
		if !assert.Equal(t, test.expected, actual) {
			t.Fatal(test.name)
		}
	}
}
