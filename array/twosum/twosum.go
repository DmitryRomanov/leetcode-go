package twosum

func twoSum(numbers []int, target int) []int {
	for i := range numbers {
		searchNumber := target - numbers[i]
		searchPos := search(searchNumber, numbers[i+1:])
		if searchPos != -1 {
			return []int{i, searchPos + i + 1}
		}
	}
	return nil
}

func search(needle int, haystack []int) int {
	for j, v := range haystack {
		if v == needle {
			return j
		}
	}
	return -1
}
