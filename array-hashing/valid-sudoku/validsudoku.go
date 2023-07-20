package validsudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		seenRow := make(map[byte]bool)
		seenColumn := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := seenRow[board[i][j]]; !ok {
					seenRow[board[i][j]] = true
				} else {
					return false
				}
			}
			if board[j][i] != '.' {
				if _, ok := seenColumn[board[j][i]]; !ok {
					seenColumn[board[j][i]] = true
				} else {
					return false
				}
			}

			if i%3 == 0 && j%3 == 0 {
				seenSubBox := make(map[byte]bool)
				for k := 0; k < 3; k++ {
					for l := 0; l < 3; l++ {
						if board[k+i][l+j] != '.' {
							if _, ok := seenSubBox[board[k+i][l+j]]; !ok {
								seenSubBox[board[k+i][l+j]] = true
							} else {
								return false
							}
						}
					}
				}
			}

		}

	}

	return true
}
