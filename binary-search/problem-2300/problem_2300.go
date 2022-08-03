package problem_2300

import (
	"math"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	result := []int{}
	sort.Ints(potions)

	for i := range spells {
		minSuccessPotion := math.Ceil(float64(success) / float64(spells[i]))
		greaterIndex := indexGreaterThanValue(potions, int(minSuccessPotion))
		if greaterIndex != -1 {
			result = append(result, len(potions)-greaterIndex)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func indexGreaterThanValue(potions []int, searchValue int) int {
	left := 0
	right := len(potions) - 1

	median := (left + right) / 2
	result := -1
	for right >= left {
		median = (left + right) / 2
		if potions[median] >= searchValue {
			right = median - 1
			result = median
		} else if potions[median] < searchValue {
			left = median + 1
		}
	}
	return result
}
