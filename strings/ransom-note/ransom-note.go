package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	var counts [26]int
	for _, v := range magazine {
		counts[int(v-'a')]++
	}
	for _, v := range ransomNote {
		counts[int(v-'a')]--
		if counts[int(v-'a')] < 0 {
			return false
		}
	}
	return true
}
