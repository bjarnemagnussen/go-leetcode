package problem0165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ret := 1
	slice1 := strings.Split(version1, ".")
	slice2 := strings.Split(version2, ".")
	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
		ret = -1
	}

	for i := range slice1 {
		d1, _ := strconv.Atoi(slice1[i])
		var d2 int
		if len(slice2) <= i {
			d2 = 0
		} else {
			d2, _ = strconv.Atoi(slice2[i])
		}

		switch {
		case d1 > d2:
			return ret
		case d1 < d2:
			return -1 * ret
		}
	}

	return 0
}
