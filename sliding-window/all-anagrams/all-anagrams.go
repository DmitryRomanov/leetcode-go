package allanagrams

import (
	"fmt"
	"sort"
)

func findAnagrams(s string, p string) []int {
	left := 0
	right := len(p)
	var result []int
	stringChars := []rune(s)
	sampleChars := RuneSlice([]rune(p))
	sort.Sort(sampleChars)

	for right < len(stringChars) {
		window := RuneSlice(stringChars[left:right])
		sort.Sort(window)
		fmt.Println()
		if sampleChars.Equal(window) {

		}
	}
	return result
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p RuneSlice) Equal(sample RuneSlice) bool { return false }
