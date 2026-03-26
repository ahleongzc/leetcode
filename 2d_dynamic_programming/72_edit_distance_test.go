package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name, word1, word2 string
		expected           int
	}{
		{
			name:     "ex 1",
			word1:    "horse",
			word2:    "ros",
			expected: 3,
		},
		{
			name:     "ex 2",
			word1:    "intention",
			word2:    "execution",
			expected: 5,
		},
	}

	for _, test := range tests {
		actual := minDistance(test.word1, test.word2)
		if !assert.Equal(t, test.expected, actual) {
			t.Fatal(test.name)
		}
	}
}
