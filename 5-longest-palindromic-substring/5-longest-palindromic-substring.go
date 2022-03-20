func longestPalindrome(s string) string {
    n := len(s)

	mat := make([][]bool, 0)

	for i := 0; i < n; i++ {
		subMat := make([]bool, n)
		mat = append(mat, subMat)
	}

	var res string

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || mat[i+1][j-1]) {
				mat[i][j] = true
			}

			if mat[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}

	return res
}