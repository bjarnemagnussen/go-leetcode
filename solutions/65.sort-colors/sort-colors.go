package problem0065

const (
	red = iota
	white
	blue
)

func sortColors(nums []int) {
	// Variable `r` denotes the index for the first item in the sorted list
	// different from the colour red.
	var r int
	// Variable `b` denotes the index for the last item in the sorted list
	// different from the colour blue.
	b := len(nums) - 1

	for cur := 0; cur <= b; {
		switch {
		case nums[cur] == blue:
			nums[cur], nums[b] = nums[b], nums[cur]
			b--
		case nums[cur] == red:
			nums[cur], nums[r] = nums[r], nums[cur]
			r++
			fallthrough
		default:
			cur++
		}
	}
}
