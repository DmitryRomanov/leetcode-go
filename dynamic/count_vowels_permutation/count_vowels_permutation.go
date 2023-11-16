// https://leetcode.com/problems/count-vowels-permutation/
package count_vowels_permutation

func countVowelPermutation(n int) int {
	rules := map[string][]string{
		"a": {"e"},
		"e": {"a", "i"},
		"i": {"a", "e", "o", "u"},
		"o": {"i", "u"},
		"u": {"a"},
	}

	dp := make([]map[string]int, 0, n)

	for i := 0; i < n; i++ {
		curDp := make(map[string]int, len(rules))
		for sym, curRules := range rules {
			if i == 0 {
				curDp[sym] = 1
			} else {
				curDp[sym] = 0
				for _, rule := range curRules {
					curDp[sym] += dp[i-1][rule]
				}
			}
		}
		dp = append(dp, curDp)
	}

	result := 0

	for _, val := range dp[n-1] {
		result += val
	}

	return result % (1e9 + 7)
}
