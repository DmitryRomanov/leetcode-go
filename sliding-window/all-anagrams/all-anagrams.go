package allanagrams

func findAnagrams(s string, p string) []int {
	left := 0
	right := len(p)
	var result []int
	stringChars := []rune(s)

	window := stringChars[left:right]
	sampleHash := calculateHash([]rune(p))
	windowHash := calculateHash(window)

	for right < len(stringChars)-1 {
		right++
		char := stringChars[right]
		j := windowHash[char]
		if j == 0 {
			windowHash[char] = 1
		} else {
			windowHash[char]++
		}

		prevChar := stringChars[left]
		windowHash[prevChar]--
		left++

		if isAnagrams(windowHash, sampleHash) {
			result = append(result, left)
		}
	}
	return result
}

func isAnagrams(window, sample map[rune]int) bool {
	for key, value := range sample {
		j := window[key]
		if j == 0 {
			window[key] = value
		} else {
			window[key] += value
		}
	}

	for _, value := range window {
		if value%2 != 0 {
			return false
		}
	}

	return true
}
func calculateHash(p []rune) map[rune]int {
	hash := make(map[rune]int)
	for _, r := range p {
		j := hash[r]
		if j == 0 {
			hash[r] = 1
		} else {
			hash[r]++
		}
	}
	return hash
}
