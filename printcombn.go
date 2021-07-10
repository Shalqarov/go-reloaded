package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	num := make([]int, n)
	for i := 0; i < n; i++ {
		num[i] = i
		z01.PrintRune(int32(num[i] + 48)) //
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	for {
		num[n-1]++
		for i := n - 2; i >= 0; i-- {
			if num[i+1] >= 10 {
				num[i]++
			}
		}
		for i := 0; i < n-1; i++ {
			if num[i+1] >= 10 {
				num[i+1] = num[i] + 1
			}
		}
		if num[n-1] >= 10 {
			continue
		}

		for i := 0; i < n; i++ {
			z01.PrintRune(int32(num[i] + 48))
		}
		if num[0] >= 10-n {
			z01.PrintRune('\n')
			break
		} else {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
