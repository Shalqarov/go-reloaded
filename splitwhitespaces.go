package student

func SplitWhiteSpaces(s string) []string {
	cnt := 0
	for i := range s {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			cnt++
		}
	}
	splitted := make([]string, 0)
	l := len(s)
	temp := ""
	for i := 0; i < l-1; i++ {
		if s[i] != ' ' || s[i] != '\t' || s[i] != '\n' {
			for j := i; j < l; j++ {
				if s[j] == ' ' || s[j] == '\t' || s[j] == '\n' {
					break
				}
				temp += string(s[j])
				i = j
			}
			if temp != "" {
				splitted = append(splitted, temp)
				temp = ""
			}
		}
	}
	return splitted
}
