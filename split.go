package student

func Split(s, sep string) []string {
	splitted := make([]string, 0)
	lS := len(s)
	lSep := len(sep)
	step := 0
	for i := 0; i+lSep < lS; i++ {
		if s[i:i+lSep] == sep {
			splitted = append(splitted, s[step:i])
			step = i + lSep
			i += lSep
		}
	}
	if sep != s[step:] {
		splitted = append(splitted, s[step:])
	}
	return splitted
}
