package count_battleships

func countBattleships(board [][]byte) (k int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' && ((i == 0 || board[i-1][j] != 'X') && (j == 0 || board[i][j-1] != 'X')) {
				k++
			}
		}
	}

	return
}
