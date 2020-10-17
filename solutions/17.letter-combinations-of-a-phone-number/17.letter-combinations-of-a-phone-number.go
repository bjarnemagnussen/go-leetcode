package problem0017

import (
	"strings"
)

var numberPad = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	input := make([][]rune, len(digits))
	for i, d := range digits {
		input[i] = numberPad[d]
	}
	return perm(input...)
}

func perm(parts ...[]rune) []string {
	res := make([]string, 0)
	for at := make([]int, len(parts)); ; at[len(at)-1]++ {
		sb := strings.Builder{}
		for i := len(at) - 1; i >= 0; i-- {
			if i != 0 && at[i] == len(parts[i]) {
				at[i] = 0
				at[i-1]++
			}
		}
		if at[0] == len(parts[0]) {
			break
		}

		// Do permutation
		for i := range parts {
			sb.WriteRune(parts[i][at[i]])
		}
		res = append(res, sb.String())
	}

	return res
}
