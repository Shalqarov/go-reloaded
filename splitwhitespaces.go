package student

func SplitWhiteSpaces(s string) []string {
	splitted := make([]string, 0)
	l := len(s)
	for i, j := 0, 0; i < l; i++ {
		if s[i] != ' ' || s[i] != '\t' || s[i] != '\n' {
			splitted[j] += string(s[i])

		} else {
			j++
		}
	}
	return splitted
}
