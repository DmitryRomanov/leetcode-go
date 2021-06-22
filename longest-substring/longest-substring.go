package longestsubstring

import "strings"

func lengthOfLongestSubstring(s string) int {
	right := 0
	result := 0
	left := 0
	stringChars := []rune(s)
	var substring []rune
	for right < len(s) {
		letter := stringChars[right]
		if strings.ContainsRune(string(substring), letter) {
			result = max(result, len(substring))
			left += strings.IndexRune(string(substring), letter) + 1
		}
		right++
		substring = stringChars[left:right]
	}

	return max(result, len(substring))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
