package subsets_2

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	var variants []int

	count := 2 << len(nums)
	for i := 0; i < int(count); i++ {
		variants = []int{}
		for j, v := range nums {
			mask := 1 << j
			if (i & mask) > 0 {
				variants = append(variants, v)
			}
		}
		sort.Ints(variants)
		if !hasDup(result, variants) {
			result = append(result, variants)
		}

	}

	return result
}

func hasDup(result [][]int, nums []int) bool {
	for _, resultNum := range result {
		if isEqual(resultNum, nums) {
			return true
		}
	}
	return false
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
