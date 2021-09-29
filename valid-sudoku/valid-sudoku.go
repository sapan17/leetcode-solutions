func isValidSudoku(board [][]byte) bool {
	// Check rows
	for row := 0; row < 9; row++ {
		panel := make([]bool, 10)
		for col := 0; col < 9; col++ {
			ch := board[row][col]
			if ch != '.' {
				num := int(ch) - int('0')
				if panel[num] {
					return false
				} else {
					panel[num] = true
				}
			}
		}
	}
	// Check columns
	for col := 0; col < 9; col++ {
		panel := make([]bool, 10)
		for row := 0; row < 9; row++ {
			ch := board[row][col]
			if ch != '.' {
				num := int(ch) - int('0')
				if panel[num] {
					return false
				} else {
					panel[num] = true
				}
			}
		}
	}
	// Check mini-squares
	for bigCol := 0; bigCol < 3; bigCol++ {
		for bigRow := 0; bigRow < 3; bigRow++ {
			panel := make([]bool, 10)
			baseCol := 3 * bigCol
			baseRow := 3 * bigRow
			for smallCol := 0; smallCol < 3; smallCol++ {
				for smallRow := 0; smallRow < 3; smallRow++ {
					ch := board[baseRow + smallRow][baseCol + smallCol]
					if ch != '.' {
						num := int(ch) - int('0')
						if panel[num] {
							return false
						} else {
							panel[num] = true
						}
					}
				}
			}
		}
	}
	return true
}