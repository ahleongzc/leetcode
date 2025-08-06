package arraysandhashing

import (
	"reflect"
	"testing"
)

func TestStringCoder_EncodeDecode(t *testing.T) {
	sc := &StringCoder{}

	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "Basic case",
			input: []string{"hello", "world"},
		},
		{
			name:  "Empty list",
			input: []string{},
		},
		{
			name:  "List with empty string",
			input: []string{""},
		},
		{
			name:  "List with multiple empty strings",
			input: []string{"", "", ""},
		},
		{
			name:  "Strings with delimiter character",
			input: []string{"a|b", "c|d|e", "||"},
		},
		{
			name:  "Strings with numbers and special characters",
			input: []string{"123", "!@#$", "test123"},
		},
		{
			name:  "Long string",
			input: []string{string(make([]byte, 1000))}, // 1000 null bytes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := sc.Encode(tt.input)
			decoded := sc.Decode(encoded)

			if !reflect.DeepEqual(tt.input, decoded) {
				t.Errorf("Decode(Encode(%v)) = %v; want %v", tt.input, decoded, tt.input)
			}
		})
	}
}
