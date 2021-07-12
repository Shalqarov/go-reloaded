package student

func Split(s, sep string) []string {
	checker := false
	splitted := make([]string, 0)
	l := len(s)
	temp := ""
	for i := 0; i < l; i++ {
		if s[i] != sep[0] {
			for j := i; j < l; j++ {
				if s[j] == sep[0] {
					for idx := 0; idx < len(sep); idx++ {
						t := j
						if s[t] != sep[idx] {
							checker = false
							break
						} else {
							checker = true
						}
					}
				}
				if checker {
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
