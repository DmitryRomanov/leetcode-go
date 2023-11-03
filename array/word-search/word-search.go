// https://leetcode.com/problems/word-search/
package word_search

func exist(board [][]byte, word string) bool {
	h := len(board)
	w := len(board[0])

	var isWordThere func(word string, index, y, x int) bool
	isWordThere = func(word string, index, y, x int) bool {
		if index == len(word)-1 {
			return word[index] == board[y][x]
		}

		if word[index] != board[y][x] {
			return false
		}

		directions := [][]int{
			{y + 1, x},
			{y - 1, x},
			{y, x + 1},
			{y, x - 1},
		}
		cachedValued := board[y][x]
		board[y][x] = ' '
		for _, direction := range directions {
			y1 := direction[0]
			x1 := direction[1]
			if y1 >= 0 && y1 < h && x1 >= 0 && x1 < w {
				if isWordThere(word, index+1, y1, x1) {
					return true
				}
			}
		}
		board[y][x] = cachedValued
		return false
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if isWordThere(word, 0, y, x) {
				return true
			}
		}
	}

	return false
}
