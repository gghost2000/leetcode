func freqAlphabets(s string) string {
    	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			val, _ := strconv.Atoi(s[i-2 : i])

			result = string(val-1+'a') + result
			i -= 2
		} else {
			result = string(int(s[i]-'0')-1+'a') + result
		}
	}

	return result
}