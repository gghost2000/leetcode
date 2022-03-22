func isSubsequence(s string, t string) bool {
    
	if len(s) == 0 {
		return true
	}
    
	n := len(t)

	location := 0

	result := false
	for i := 0; i < n; i++ {
		if s[location] == t[i] {
			location++
		}

		if location == len(s) {
			result = true
			break
		}

	}

	return result
}