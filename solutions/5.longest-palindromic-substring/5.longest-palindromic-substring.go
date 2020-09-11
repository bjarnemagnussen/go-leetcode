package problem0005

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	begin, maxLen := 0, 1
	for i := 0; i < len(s)-1; i++ {
		cut := 1
		for k := 0; k < (len(s)-cut-i+1)/2; k++ {
			if s[i+k] != s[len(s)-cut-k] {
				cut++
				k = -1
			}
		}
		if len(s)-cut+1-i > maxLen-begin {
			begin, maxLen = i, len(s)-cut+1
		}
	}
	return s[begin:maxLen]
}
