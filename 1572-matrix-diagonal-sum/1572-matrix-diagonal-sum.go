func diagonalSum(mat [][]int) int {
    sum := 0

	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
	}

	for i := len(mat); i > 0; i-- {
		sum += mat[len(mat)-i][i-1]
	}

	if len(mat)%2 == 1 {
		sum -= mat[len(mat)/2][len(mat)/2]
	}

	return sum
}