package problem0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := make([]int, 0)
	for ; x > 0; x /= 10 {
		digits = append(digits, x%10)
	}

	for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
		if digits[left] != digits[right] {
			return false
		}
	}

	return true
}
