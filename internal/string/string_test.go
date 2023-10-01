package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitByOperators(t *testing.T) {
	tests := []struct {
		name string
		s    string
		expS []string
	}{
		{
			name: "Empty string",
			s:    "",
			expS: nil,
		},
		{
			name: "Only spaces",
			s:    "    ",
			expS: nil,
		},
		{
			name: "Only integer",
			s:    "12",
			expS: []string{"12"},
		},
		{
			name: "Only operator",
			s:    "+",
			expS: []string{"+"},
		},
		{
			name: "Multiple operators",
			s:    "+/-",
			expS: []string{"+", "/", "-"},
		},
		{
			name: "Numbers and one operator",
			s:    "1+32",
			expS: []string{"1", "+", "32"},
		},
		{
			name: "Numbers, operators, and spaces",
			s:    "   + 23   - 32 / 1 *",
			expS: []string{"+", "23", "-", "32", "/", "1", "*"},
		},
		{
			name: "Numbers, operators, spaces, and invalid operators",
			s:    "   + 23   - 32 / 1 * !",
			expS: []string{"+", "23", "-", "32", "/", "1", "*", "!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := SplitByOperators(tt.s)
			assert.Equal(t, tt.expS, ss)
		})
	}
}
