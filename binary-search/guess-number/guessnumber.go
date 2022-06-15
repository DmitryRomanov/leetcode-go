package guessnumber

import "math"

var pick int

func guessNumber(n int) int {
	left := 0
	right := n
	for left <= right {
		median := int(math.Round(float64((left + right) / 2)))
		guessResult := guess(median)
		if guessResult == 0 {
			return median
		} else if guessResult == -1 {
			right = median - 1
		} else if guessResult == 1 {
			left = median + 1
		}
	}

	return 0
}

func guess(n int) int {
	if n > pick {
		return -1
	} else if n < pick {
		return 1
	}
	return 0
}
