package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	var counts [26]int
	for _, v := range magazine {
		counts[v-'a']++
	}
	for _, v := range ransomNote {
		counts[v-'a']--
		if counts[v-'a'] < 0 {
			return false
		}
	}
	return true
}
