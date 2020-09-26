package problem0011

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	maxArea := 0

	for start < end {
		if height[start] < height[end] {
			maxArea = max(maxArea, (end-start)*height[start])
			start++
		} else {
			maxArea = max(maxArea, (end-start)*height[end])
			end--
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
