func isHappy(n int) bool {
    maps := make(map[int]bool, 0)
    result := true

	for n != 1 {
		val := 0

		for n != 0 {
			sum := n % 10
			val += sum * sum
			n /= 10
		}

		if maps[val] {
			result = false
			break
		}

		maps[val] = true

		n = val

	}

	return result

}