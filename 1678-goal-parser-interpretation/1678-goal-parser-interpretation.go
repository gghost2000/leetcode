func interpret(command string) string {
    Go := "G"
	O := "()"
	al := "(al)"

	result := ""

	for i := 0; i < len(command); {
		if strings.HasPrefix(command[i:], Go) {
			result += Go
			i += len(Go)
		}

		if strings.HasPrefix(command[i:], O) {
			result += "o"
			i += len(O)
		}

		if strings.HasPrefix(command[i:], al) {
			result += "al"
			i += len(al)
		}
	}

	return result
}