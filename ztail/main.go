package main

import "os"

func main() {
	args := os.Args[1:]
	argsSize := len(args)
	if argsSize < 3 {
		os.Exit(0)
	} else {
		Carg := false
		for i := 0; i < argsSize; i++ {
			if args[i] == "-c" {
				Carg = true
				if i+1 == argsSize {
					os.Exit(0)
				}
			}
			if !Carg {
				os.Exit(0)
			}
		}
	}
}

func toInt(s string) int {
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
			num = num*10 + int(rune(s[i]-48))
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
