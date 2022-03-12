func wordBreak(s string, wordDict []string) bool {
    dp := make(map[int]bool, 0)
	dp[0] = true
	max := 0
	size := len(s)
    
    	for i := 0; i < size; i++ {
		if dp[i] {
			for _, word := range wordDict {
				if strings.HasPrefix(s[i:], word) {
					point := i + len(word)
					dp[point] = true
					max = findmax(max, point)
				}
			}
		}

	}
    
    return dp[size]
    
}


func findmax(a, b int) int {
	if a > b {
		return a
	}

	return b
}