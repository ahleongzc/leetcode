package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "example 1",
			s:        "XYYX",
			k:        2,
			expected: 4,
		},
		{
			name:     "example 2",
			s:        "AAABABB",
			k:        1,
			expected: 5,
		},
		{
			name:     "example 3",
			s:        "ABAA",
			k:        0,
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := characterReplacement(test.s, test.k)
			assert.Equal(t, test.expected, result, "%v: expected %v, got %v", test.name, test.expected, result)
		})
	}
}
