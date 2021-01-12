package problem0007

func reverse(x int) int {
	var ans int
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}

	if ans > 2147483647 || ans < -2147483648 {
		return 0
	}

	return ans
}
