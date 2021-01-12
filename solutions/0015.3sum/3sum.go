package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	unique := make([]int, 0)
	counter := make(map[int]int)
	for _, n := range nums {
		if counter[n] == 0 {
			unique = append(unique, n)
		}
		counter[n]++
	}
	sort.Ints(unique)

	// Find x, y, z in nums, s.t. x+y+z=0.
	// Fix x.
	for i, x := range unique {
		if 3*x == 0 && counter[x] >= 3 {
			res = append(res, []int{x, x, x})
		}

		// Fix y.
		for j := i + 1; j < len(unique); j++ {
			y := unique[j]

			if 2*x+y == 0 && counter[x] >= 2 {
				res = append(res, []int{x, x, y})
			}

			if x+2*y == 0 && counter[y] >= 2 {
				res = append(res, []int{x, y, y})
			}

			// Find z.
			z := -x - y
			if z > x && z > y && counter[z] > 0 {
				res = append(res, []int{x, y, z})
			}
		}
	}
	return res
}
