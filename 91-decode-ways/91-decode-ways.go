func numDecodings(s string) int {
    dp := make([]int, 101)
	dp[0] = 1
	dp[1] = 1
	if s[0] == '0' {
		return 0
	}

	size := len(s)

	for i := 2; i <= size; i++ {
		one, _ := strconv.Atoi(string(s[i-1]))
		two, _ := strconv.Atoi(string(s[i-2 : i]))

		if one != 0 {
			dp[i] += dp[i-1]
		}

		if two >= 10 && two <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[size]
}