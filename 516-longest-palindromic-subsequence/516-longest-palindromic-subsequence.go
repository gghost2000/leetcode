func longestPalindromeSubseq(s string) int {
    n := len(s)

	mat := make([][]int, 0)

	for i := 0; i < n; i++ {
		subMat := make([]int, n)
		mat = append(mat, subMat)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				mat[i][j] = 1
			} else if s[i] == s[j] {
				mat[i][j] = mat[i+1][j-1] + 2
			} else {
				mat[i][j] = max(mat[i+1][j], mat[i][j-1])
			}
		}
	}

	return mat[0][n-1]
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}