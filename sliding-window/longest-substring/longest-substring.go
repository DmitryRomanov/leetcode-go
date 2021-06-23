package longestsubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	right := 0
	left := 0
	stringChars := []rune(s)
	hash := make(map[rune]bool)
	result := 1

	for right < len(s) {
		rightLetter := stringChars[right]
		leftLetter := stringChars[left]
		if !hash[rightLetter] {
			result = max(result, right-left+1)
			hash[rightLetter] = true
			right++
		} else {
			hash[leftLetter] = false
			left++
		}
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
