package maxpointcards

func maxScore(cardPoints []int, k int) int {
	totalPts := sum(cardPoints)
	left := 0
	right := len(cardPoints) - k - 1

	result := 0
	temporary := sum(cardPoints[left : right+1])
	for {
		result = max(totalPts-temporary, result)
		if left == len(cardPoints) {
			break
		}
		temporary -= cardPoints[left]
		left++
		right++
		if right == len(cardPoints) {
			break
		}
		temporary += cardPoints[right]
	}
	return result
}

func sum(array []int) int {
	result := 0
	for i := range array {
		result += array[i]
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
