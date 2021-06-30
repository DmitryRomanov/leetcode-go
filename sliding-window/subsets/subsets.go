package subsets

import (
	"math"
	"sort"
)

func subsets(nums []int) [][]int {
	var result [][]int
	var variants []int

	count := math.Pow(2, float64(len(nums)))
	for i := 0; i < int(count); i++ {
		variants = []int{}
		for j, v := range nums {
			mask := 1 << j
			if (i & mask) != 0 {
				variants = append(variants, v)
			}
		}
		sort.Ints(variants)
		result = append(result, variants)
	}

	return result
}
