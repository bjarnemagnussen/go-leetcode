package problem0027

func removeElement(nums []int, val int) int {
	tot := len(nums)
	if tot == 0 {
		return 0
	}

	p := tot - 1
	for nums[p] == val {
		p--
		if p < 0 {
			return 0
		}
	}

	for i := 0; i < tot && i <= p; i++ {
		if nums[i] == val {
			nums[i], nums[p] = nums[p], nums[i]
			for nums[p] == val {
				p--
			}
		}
	}

	return p + 1
}
