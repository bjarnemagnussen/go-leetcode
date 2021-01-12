package problem0674

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	var longest, cut int
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i-1] >= nums[i] {
			longest = max(longest, i-cut)
			cut = i
		}
	}
	return longest
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
