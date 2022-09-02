package combination_sum

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	maxElements := target / findMin(candidates)
	variants := combrep(int(maxElements), candidates)

	fmt.Println(variants)

	for i := range variants {
		variant := variants[i]
		if check(variant, target) {
			result = append(result, variant)
		}
	}

	return result
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func combrep(n int, lst []int) [][]int {
	if n == 0 {
		return [][]int{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	r := combrep(n, lst[1:])
	for _, x := range combrep(n-1, lst) {
		r = append(r, append(x, lst[0]))
	}
	return r
}

func check(lst []int, target int) bool {
	sum := 0
	for i := range lst {
		sum += lst[i]
	}
	return sum == target
}

func findMin(a []int) (min int) {
	min = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}
