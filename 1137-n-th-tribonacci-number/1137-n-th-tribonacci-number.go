func tribonacci(n int) int {
    var dp [41]int
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 0; i <= n-3; i++ {
		dp[i+3] = dp[i] + dp[i+1] + dp[i+2]
	}
    
    return dp[n]
}