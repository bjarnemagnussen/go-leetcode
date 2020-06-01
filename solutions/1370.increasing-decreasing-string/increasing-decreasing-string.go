package problem1370

import "strings"

func sortString(s string) string {
	var counter [26]rune
	var total int
	for _, ch := range s {
		counter[ch-'a']++
		total++
	}

	var sb strings.Builder
	for i := 0; total > 0; i++ {
		idx := i % (2 * 26)
		if idx > 25 {
			idx = 25 - (i % 26)
		}

		if val := counter[idx]; val > 0 {
			sb.WriteRune(rune(idx) + 'a')
			counter[idx]--
			total--
		}
	}
	return sb.String()
}
