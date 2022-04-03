func searchMatrix(matrix [][]int, target int) bool {
    result := false

	if matrix[0][0] > target || matrix[len(matrix)-1][len(matrix[0])-1] < target {
		return result
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				result = true
				break
			}
		}
	}

	return result
}