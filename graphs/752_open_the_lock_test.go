package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenLock(t *testing.T) {
	tests := []struct {
		name     string
		deadends []string
		target   string
		expected int
	}{
		{
			name:     "example 0",
			deadends: []string{},
			target:   "0001",
			expected: 1,
		},
		{
			name: "example 1",
			deadends: []string{
				"0201",
				"0101",
				"0102",
				"1212",
				"2002",
			},
			target:   "0202",
			expected: 6,
		},
		{
			name: "example 2",
			deadends: []string{
				"8888",
			},
			target:   "0009",
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := openLock(test.deadends, test.target)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf("test case %v - expected: %v, got %v", test.name, test.expected, output)
			}
		})
	}
}
