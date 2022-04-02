func areAlmostEqual(s1 string, s2 string) bool {
    if s1 == s2 {
		return true
	}

	result := false

	list := make([]int, 0)

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			list = append(list, i)
		}
	}

	if len(list) == 2 && s1[list[0]] == s2[list[1]] && s1[list[1]] == s2[list[0]] {
		result = true
	}

        return result
}