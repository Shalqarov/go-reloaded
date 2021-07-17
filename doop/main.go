package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	//fmt.Println(args) // проверка на правильность аргументы ( нужно удалить )
	argsLen := len(args)
	if argsLen != 3 {
		z01.PrintRune('0') // если введут меньше 3 или больше аргументов
		z01.PrintRune('\n')
	} else {
		if isNum(args[0]) && isNum(args[2]) && isOp(args[1]) { // проверяю цифру ли ввели
			var a, b int64
			a = Atoi64(args[0])
			b = Atoi64(args[2])
			op := args[1]
			if a == 0 && args[0] != "0" || b == 0 && args[2] != "0" {
				z01.PrintRune('0')
			} else {
				switch op {
				case "+":
					{
						if a > 0 && b > 0 && a+b < 0 {
							z01.PrintRune('0')
						} else {
							res := itoa64(a + b)
							lRes := len(res)
							for i := 0; i < lRes; i++ {
								z01.PrintRune(rune(res[i]))
							}
						}
					}
				case "-":
					{
						if a < 0 && b < 0 && a-b > 0 {
							z01.PrintRune('0')
						} else {
							res := itoa64(a - b)
							lRes := len(res)
							for i := 0; i < lRes; i++ {
								z01.PrintRune(rune(res[i]))
							}
						}
					}
				case "*":
					{
						if a == 0 || b == 0 {
							z01.PrintRune('0')
						} else {
							res := itoa64(a * b)
							lRes := len(res)
							if (a > 0 && b > 0) && (a*b < a || a*b < b) {
								z01.PrintRune('0')
							} else if (a > 0 && b < 0) && (a*b < a || a*b < b) || (a < 0 && b > 0) && (a*b < a || a*b < b) {
								z01.PrintRune('0')
							} else {
								for i := 0; i < lRes; i++ {
									z01.PrintRune(rune(res[i]))
								}
							}
						}
					}
				case "/":
					{
						if b == 0 {
							err := "No division by 0"
							errLen := len(err)
							for i := 0; i < errLen; i++ {
								z01.PrintRune(rune(err[i]))
							}
						} else {
							res := itoa64(a / b)
							lRes := len(res)
							for i := 0; i < lRes; i++ {
								z01.PrintRune(rune(res[i]))
							}
						}
					}
				case "%":
					{
						if b == 0 {
							err := "No modulo by 0"
							errLen := len(err)
							for i := 0; i < errLen; i++ {
								z01.PrintRune(rune(err[i]))
							}
						} else {
							res := itoa64(a % b)
							lRes := len(res)
							for i := 0; i < lRes; i++ {
								z01.PrintRune(rune(res[i]))
							}
						}
					}
				}
			}
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('0') // проверка на вход ( нужно удалить )
			z01.PrintRune('\n')
		}
	}
}

func isOp(s string) bool {
	lS := len(s)
	if lS != 1 {
		return false
	}
	if s[0] == '+' || s[0] == '-' || s[0] == '/' || s[0] == '*' || s[0] == '%' {
		return true
	}
	return false
}

func isNum(s string) bool {
	lS := len(s)
	start := 0
	if s[0] == '-' {
		start = 1
	}
	for i := start; i < lS; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func itoa64(n int64) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	s := []rune{}
	if n < 0 {
		isNegative = true
		n *= -1
	}
	for i := 0; n > 0; i++ {
		s = append(s, rune(n%10+'0'))
		n /= 10
	}
	if isNegative {
		s = append(s, '-')
	}
	res := reverse(string(s))
	return res
}

func reverse(s string) string {
	lS := len(s)
	res := []rune{}

	for i := lS - 1; i >= 0; i-- {
		res = append(res, rune(s[i]))
	}
	return string(res)
}

func Atoi64(s string) int64 {
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
		var num int64 = 0
		for i := startIndex; i < len(s); i++ {
			num = num*10 + int64(rune(s[i]-48))
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
