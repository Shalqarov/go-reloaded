package student

func Split(s, sep string) []string {
	splitted := make([]string, 0)
	lS := len(s)
	lSep := len(sep)
	step := 0
	if sep == "" {
		var char rune
		for i := 0; i < lS; i++ {
			char = rune(s[i])
			splitted = append(splitted, string(char))
		}
		return splitted
	}
	for i := 0; i+lSep <= lS; i++ {
		if s[i:i+lSep] == sep {
			splitted = append(splitted, s[step:i])
			step = i + lSep
			i += lSep - 1
		}
	}
	if sep != s[step:] {
		splitted = append(splitted, s[step:])
	}
	return splitted
}
