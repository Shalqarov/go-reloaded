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
		fmt.Println(str)
		vowels := map[rune]bool{'a': true, 'o': true,
			'u': true, 'e': true, 'i': true,
			'A': true, 'O': true, 'U': true,
			'E': true, 'I': true}
		rotatevowels := []rune{}
		idx := []int{}
		for i, symbol := range str {
			if _, isHave := vowels[symbol]; isHave {
				rotatevowels = append(rotatevowels, rune(str[i]))
				idx = append(idx, i)
			}
		}
		rtLen := len(rotatevowels)
		for i, j := 0, rtLen-1; i < j; i++ {
			t := rotatevowels[j]
			rotatevowels[j] = rotatevowels[i]
			rotatevowels[i] = t
			j--
		}
		res := ""
		for i, j := 0, 0; i < rtLen; i++ {
			if i == idx[j] {
				res += string(rotatevowels[j])
				j++
			} else {
				res += str
			}
		}
		resLen := len(res)
		for i := 0; i < resLen; i++ {
			z01.PrintRune(rune(res[i]))
		}
	}
}
