func firstUniqChar(s string) int {
    maps := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		maps[s[i]] += 1
	}

	index := -1

	for i := 0; i < len(s); i++ {
		if maps[s[i]] == 1 {
			index = i
			break
		}
	}

	return index
}