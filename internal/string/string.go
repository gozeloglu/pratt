package string

import "strings"

var operators = []uint8{'+', '-', '*', '/', '<'}

// SplitByOperators splits the given string by operators(+,-,*,/,<).
func SplitByOperators(s string) []string {
	var res []string

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return res
	}

	var a strings.Builder
	for i := 0; i < len(s); i++ {
		if isOperator(s[i]) {
			if a.Len() > 0 {
				res = append(res, strings.TrimSpace(a.String()))
				a.Reset()
			}
			a.WriteByte(s[i])
			res = append(res, a.String())
			a.Reset()
		} else {
			a.WriteByte(s[i])
		}
	}
	if a.Len() > 0 {
		res = append(res, strings.TrimSpace(a.String()))
	}
	return res
}

func isOperator(ch uint8) bool {
	for _, op := range operators {
		if op == ch {
			return true
		}
	}

	return false
}
