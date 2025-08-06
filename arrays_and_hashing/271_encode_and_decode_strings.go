package arraysandhashing

import (
	"fmt"
	"strconv"
	"strings"
)

type StringCoder struct{}

func (s *StringCoder) Encode(list []string) string {
	output := strings.Builder{}
	for _, s := range list {
		output.WriteString(fmt.Sprintf("%d|%s", len(s), s))
	}
	return output.String()
}

func (s *StringCoder) Decode(str string) []string {
	ptr := 0
	output := make([]string, 0)

	for ptr != len(str) {
		count := strings.Builder{}
		for string(str[ptr]) != "|" {
			count.WriteString(string(str[ptr]))
			ptr++
		}

		// To move past the | char
		ptr++

		countInt, _ := strconv.Atoi(count.String())
		intermediary := str[ptr : countInt+ptr]

		output = append(output, intermediary)
		ptr += countInt
		count.Reset()
	}
	return output
}
