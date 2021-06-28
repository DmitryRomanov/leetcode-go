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
		if isAnagrams(windowHash, sampleHash) {
			result = append(result, left)
		}
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

	}
	return result
}

func isAnagrams(window, sample map[rune]int) bool {
	result := make(map[rune]int)
	for key, value := range sample {
		j := result[key]
		if j == 0 {
			result[key] = value
		} else {
			result[key] += value
		}
	}

	for key, value := range window {
		j := result[key]
		if j == 0 {
			result[key] = value
		} else {
			result[key] += value
		}
	}

	for _, value := range result {
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
