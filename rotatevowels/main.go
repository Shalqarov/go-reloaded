package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	argsLen := len(args)
	if argsLen < 1 {
		z01.PrintRune('\n')
	} else {
		str := []rune{}
		for i := 0; i < argsLen; i++ {
			for j := 0; j < len(args[i]); j++ {
				str = append(str, rune(args[i][j]))
			}
			str = append(str, ' ')
		}
		strLen := len(str)
		for L, R := 0, strLen-1; L <= R; {
			if !isVowel(str[L]) {
				L++
			}
			if !isVowel(str[R]) {
				R--
			}
			if isVowel(str[L]) && isVowel(str[R]) {
				str[L], str[R] = str[R], str[L]
				L++
				R--
			}
		}
		for i := 0; i < len(str); i++ {
			z01.PrintRune(str[i])
		}
		z01.PrintRune('\n')
	}
}

func isVowel(c rune) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}
