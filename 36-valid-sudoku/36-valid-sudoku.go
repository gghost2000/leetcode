func isValidSudoku(board [][]byte) bool {
    arr := make([][]int, 0)

	for i := 0; i < 9; i++ {
		newArr := make([]int, 9)
		arr = append(arr, newArr)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]

			if val == '.' {
				arr[i][j] = 0
			} else {
				arr[i][j], _ = strconv.Atoi(string(val))
			}
		}
	}

	if !checkLine(arr) {
		return false
	}

	if !checkMap(arr) {
		return false
	}

    return true 
}

func checkLine(arr [][]int) bool {

	for i := 0; i < 9; i++ {
		maps := make(map[int]int, 0)
		for j := 0; j < 9; j++ {
			if arr[i][j] != 0 {
				if maps[arr[i][j]] != 0 {
					return false
				} else {
					maps[arr[i][j]] += 1
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		maps := make(map[int]int, 0)
		for j := 0; j < 9; j++ {
			if arr[j][i] != 0 {
				if maps[arr[j][i]] != 0 {
					return false
				} else {
					maps[arr[j][i]] += 1
				}
			}
		}
	}

	return true
}

func checkMap(arr [][]int) bool {
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			maps := make(map[int]int, 0)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					xVal := (i * 3) + x
					yVal := (j * 3) + y
					if arr[xVal][yVal] != 0 {
						if maps[arr[xVal][yVal]] != 0 {
							return false
						} else {
							maps[arr[xVal][yVal]] += 1
						}
					}

				}

			}

		}
	}

	return true
}
