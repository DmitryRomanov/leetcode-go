package count_battleships

func countBattleships(board [][]byte) (k int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == 'X' {
				board[row][col] = '.'
				k++
				if (col+1) < len(board[row]) && board[row][col+1] == 'X' {
					// go right
					col++
					for (col) < len(board[row]) && board[row][col] == 'X' {
						board[row][col] = '.'
						col++
					}
				} else if (row+1) < len(board) && board[row+1][col] == 'X' {
					// go down
					row2 := row + 1
					for row2 < len(board) && board[row2][col] == 'X' {
						board[row2][col] = '.'
						row2++
					}
				} else {
					col++
				}
			}
		}
	}

	return
}
