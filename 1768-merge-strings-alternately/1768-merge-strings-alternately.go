func mergeAlternately(word1 string, word2 string) string {
    result := ""

	maxLen := max(len(word1), len(word2))

	for i := 0; i < maxLen; i++ {
		if i > len(word1)-1 {
			result += word2[i:]
			break
		}

		if i > len(word2)-1 {
			result += word1[i:]
			break
		}

		result += word1[i:i+1] + word2[i:i+1]
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}