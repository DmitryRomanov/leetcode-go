// https://leetcode.com/problems/valid-palindrome/description/
package valid_palindrome

import (
	"unicode"
)

func isLetterOrDigit(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left <= right {
		charLeft := unicode.ToLower(rune(s[left]))
		if !isLetterOrDigit(charLeft) {
			left++
			continue
		}
		charRight := unicode.ToLower(rune(s[right]))
		if !isLetterOrDigit(charRight) {
			right--
			continue

		}

		if charLeft != charRight {
			return false
		}

		left++
		right--
	}
	return true
}
