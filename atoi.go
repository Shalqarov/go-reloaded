package student

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	if s[0] == '+' || s[0] == '-' || s[0] >= '0' && s[0] <= '9' {
		startIndex := 0
		if s[0] == '+' || s[0] == '-' {
			startIndex = 1
		}
		for i := startIndex; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return 0
			}
		}
		num := 0
		for i := startIndex; i < len(s); i++ {
			num = num*10 + int(s[i]-48)
		}
		isNegative := false
		if s[0] == '-' {
			isNegative = true
			num *= -1
		}
		if num > 0 && isNegative == true {
			return 0
		} else if num < 0 && isNegative == false {
			return 0
		}
		return num
	}
	return 0
}
