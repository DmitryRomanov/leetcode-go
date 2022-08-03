package problem_2300

import (
	"math"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	result := make([]int, len(spells))
	sort.Ints(potions)

	for i, spell := range spells {
		minSuccessPotion := math.Ceil(float64(success) / float64(spell))
		greaterIndex := indexGreaterThanValue(potions, int(minSuccessPotion))
		if greaterIndex != -1 {
			result[i] = len(potions) - greaterIndex
		} else {
			result[i] = 0
		}
	}
	return result
}

func indexGreaterThanValue(potions []int, searchValue int) int {
	left := 0
	right := len(potions) - 1

	median := left + (right-left)>>1
	result := -1
	for right >= left {
		median = left + (right-left)>>1
		if potions[median] >= searchValue {
			right = median - 1
			result = median
		} else if potions[median] < searchValue {
			left = median + 1
		}
	}
	return result
}
