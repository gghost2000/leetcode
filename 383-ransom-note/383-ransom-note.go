func canConstruct(ransomNote string, magazine string) bool {
    	maps := make(map[byte]int, 0)

	for i := 0; i < len(magazine); i++ {
		maps[magazine[i]] += 1
	}

	result := true

	for i := 0; i < len(ransomNote); i++ {
		maps[ransomNote[i]] -= 1

		if maps[ransomNote[i]] < 0 {
			result = false
			break
		}
	}

	return result
}