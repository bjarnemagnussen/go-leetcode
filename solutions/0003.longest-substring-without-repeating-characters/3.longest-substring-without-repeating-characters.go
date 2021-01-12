package problem0003

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	lastCut := make(map[rune]int, 0)
	var total, curCut int
	for i, ch := range s {
		curCut = max(curCut, lastCut[ch])
		lastCut[ch] = i + 1
		total = max(i-curCut+1, total)
	}
	return total
}
