package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	argsLen := len(args)
	if argsLen < 1 {
		z01.PrintRune('\n')
	} else {
		str := args[0]
		for i := 1; i < argsLen; i++ {
			str += " " + args[i]
		}
		strLen := len(str)
		res := make([]rune, strLen)
		leftvowel := false
		rightvowel := false
		for l, r := 0, strLen-1; l < r; {
			if str[l] == 'a' || str[l] == 'e' || str[l] == 'i' || str[l] == 'o' || str[l] == 'u' {
				leftvowel = true
				if rightvowel == true {
					res[l] = rune(str[r])
					res[r] = rune(str[l])
					leftvowel = false
					rightvowel = false
					l++
					r++
					continue
				}
			}
			if str[r] == 'a' || str[r] == 'e' || str[r] == 'i' || str[r] == 'o' || str[r] == 'u' {
				rightvowel = true
				if leftvowel == true {
					res[l] = rune(str[r])
					res[r] = rune(str[l])
					leftvowel = false
					rightvowel = false
					l++
					r++
					continue
				}
			}
			res[l] = rune(str[l])
			res[r] = rune(str[r])
			l++
			r++
		}
		fmt.Println(res)
	}
}
