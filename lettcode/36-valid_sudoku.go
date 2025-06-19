import (
	"strconv"
)

func isValidSudoku(board [][]byte) bool {

	mapLine := make(map[string]bool)
	mapCol := make(map[string]bool)
	mapQuad := make(map[string]bool)

	//check lines

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			keyLine := strconv.Itoa(i) + "_" + string(board[i][j])
			keyCol := strconv.Itoa(j) + "_" + string(board[i][j])
			keyQuad := strconv.Itoa(i/3) + "_" + strconv.Itoa(j/3) + "_" + string(board[i][j])

			_, ok_X := mapLine[keyLine]
			_, ok_Y := mapCol[keyCol]
			_, ok_Q := mapQuad[keyQuad]

			if ok_X {
				return false
			}

			if ok_Y {
				return false
			}

			if ok_Q {
				return false
			}

			mapLine[keyLine] = true
			mapCol[keyCol] = true
			mapQuad[keyQuad] = true
		}
	}

	return true
}