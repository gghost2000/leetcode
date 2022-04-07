func isValid(s string) bool {
	list := make([]byte, 0)

	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			list = append(list, s[i])
		} else {
			if len(list) == 0 {
				return false
			}

			val := list[len(list)-1]
			list = list[0 : len(list)-1]

			if s[i] == ')' {
				if val != '(' {
					return false
				}
			} else if s[i] == ']' {
				if val != '[' {
					return false
				}
			} else {
				if val != '{' {
					return false
				}
			}

		}
	}

	if len(list) != 0 {
		return false
	}

	return true
}