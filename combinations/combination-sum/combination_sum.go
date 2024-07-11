package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	for _, candidate := range candidates {
		if target-candidate >= 0 {
			combinations := combinationSum(candidates, target-candidate)
			result = append(result, combinations...)
		}
	}
	return result
}
