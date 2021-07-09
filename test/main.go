package main
//import "fmt"
import "github.com/01-edu/z01"

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	if s[0] == '+' || s[0] == '-' || s[0] >='0' && s[0] <= '9'{
		startIndex:=0
		if s[0] == '+' || s[0] == '-' {
			startIndex = 1
		}
		for i:=startIndex; i < len(s); i++{
			if s[i] < '0' || s[i] > '9' {
				return 0
			}
		}
		num := 0
		for i:=startIndex;i<len(s);i++{
			num = num*10 + int(rune(s[i] - 48))
			if num < 0 {
				break
			}
		}
		isNegative:=false
		if s[0] == '-' {
			isNegative = true
			num*=-1
		}
		if num > 0 && isNegative == true {
			return 0
		} else if num < 0 && isNegative ==false {
			return 0
		}
		return num
	} else {
		return 0
	}
}
func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}else if power == 0 {
		return 1
	}else if power == 1 {
		return nb
	}
	return nb * RecursivePower(nb, power - 1)
}
func PrintCombN(n int) {
	num := make([]int,n)
	for i := 0; i < n; i++ {
		num[i] = i
		z01.PrintRune(int32(num[i] + 48))
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	for {
		num[n-1] += 1
		for i:=n - 2; i >= 0; i-- {
			if num[i+1] >= 10 {
				num[i] += 1
			}
		}
		for i:=0; i < n -1; i++ {
			if num[i + 1] >=10 {
				num[i + 1] = num[i] + 1
			}
		}
		if num[n-1] >= 10 {
			continue
		}
		
		for i := 0; i < n; i++ {
			z01.PrintRune(int32(num[i] + 48))
		}
		if num[0]>=10-n{
			z01.PrintRune('\n')
			break
		}else {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		
	}

}
func PrintNbrBase(nbr int, base string) {
	
}

// func ActiveBits(n int) uint {
// }

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
