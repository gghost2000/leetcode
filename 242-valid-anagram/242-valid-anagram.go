func isAnagram(s string, t string) bool {

	sMap := getMap(s)
	tMap := getMap(t)

	result := true

	if len(s) >= len(t) {
		result = checkMap(sMap, tMap)
	} else {
		result = checkMap(tMap, sMap)
	}

	return result
}


func getMap(text string) map[byte]int {
	maps := make(map[byte]int, 0)

	for i := 0; i < len(text); i++ {
		maps[text[i]] += 1
	}

	return maps
}

func checkMap(a, b map[byte]int) bool {
	for key, _ := range a {
		if a[key] != b[key] {
			return false
		}
	}

	return true
}
