func findTheDifference(s string, t string) byte {
    sMap := makeMap(s)
	tMap := makeMap(t)

	var result byte

	for k, _ := range tMap {
		if tMap[k] > sMap[k] {
			result = k
			break
		}
	}

	return result
}

func makeMap(text string) map[byte]int {
	maps := make(map[byte]int, 0)

	for i := 0; i < len(text); i++ {
		maps[text[i]] += 1
	}

	return maps
}