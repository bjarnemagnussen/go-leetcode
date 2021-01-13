package problem0049

import "sort"

func groupAnagrams(strs []string) [][]string {
	uniques := make(map[string][]string)
	for _, word := range strs {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		uniques[string(runes)] = append(uniques[string(runes)], word)
	}

	groups := [][]string{}
	for _, group := range uniques {
		groups = append(groups, group)
	}

	return groups

	// Alternative solution:
	// We can use an array to hold unique representations of words as key
	// to the uniques map.
	//
	// uniques := make(map[[26]rune][]string)
	// for _, word := range strs {
	// 	count := [26]rune{}
	// 	for _, ch := range word {
	// 		count[ch-'a']++
	// 	}
	// 	uniques[count] = append(uniques[count], word)
	// }

	// groups := [][]string{}
	// for _, v := range uniques {
	// 	groups = append(groups, v)
	// }
	// return groups
}
