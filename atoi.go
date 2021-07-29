package student

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	if s[0] == '+' || s[0] == '-' || s[0] >= '0' && s[0] <= '9' {
		startIndex := 0
		var num = 0
		isNegative := false
		if s[0] == '-' {
			isNegative = true
		}
		if s[0] == '+' || s[0] == '-' {
			startIndex = 1
		}
		zerocount := false
		cnt := 0
		for i := startIndex; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return 0
			}
			if s[i] > '0' && s[i] <= '9' {
				zerocount = true
			}
			if zerocount {
				cnt++
			}
		}
		if cnt > 19 && !isNegative || cnt > 20 && isNegative {
			return 0
		}
		if isNegative == false {
			for i := startIndex; i < len(s); i++ {
				num = num*10 + int(rune(s[i]-48))
			}
		} else if isNegative == true {
			for i := startIndex; i < len(s); i++ {
				num = num*10 - int(rune(s[i]-48))
			}
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
