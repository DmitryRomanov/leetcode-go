// https://leetcode.com/problems/longest-palindrome/description/
package longest_palindrome

func longestPalindrome(s string) int {
	var counts [58]int
	for _, v := range s {
		counts[int(v-'A')]++
	}
	result := 0
	for _, v := range counts {
		if v == 0 {
			continue
		}
		result += v / 2 * 2
		if result%2 == 0 && v%2 == 1 {
			result++
		}
	}
	return result
}
