// https://leetcode.com/problems/longest-palindrome/description/
package longest_palindrome

func longestPalindrome(s string) int {
	var seen [58]bool
	result := 0
	for _, v := range s {
		if seen[int(v-'A')] {
			result += 2
			seen[int(v-'A')] = false
			continue
		}
		seen[int(v-'A')] = true
	}
	if len(s) > result {
		result++
	}
	return result
}
