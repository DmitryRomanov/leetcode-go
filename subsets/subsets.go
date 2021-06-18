package subsets

import "sort"

func subsets(nums []int) [][]int {
	result := [][]int{{}}

	for _, v := range nums {
		for _, j := range result {
			tempResult := append(j, v)
			sort.Ints(tempResult)
			result = append(result, tempResult)
		}
	}

	return result
}
