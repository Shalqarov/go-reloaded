package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	baseLen := len(base)
	for i := 0; i < baseLen; i++ {
		if base[i] == '-' || base[i] == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	dic := make(map[rune]bool, 0)
	for _, str := range base {
		if dic[str] {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		dic[str] = true
	}

	isNegative := false
	if nbr < 0 {
		isNegative = true
	}
	idx := nbr % baseLen
	nbr /= baseLen
	if isNegative {
		idx *= -1
		nbr *= -1
	}
	res := []int{}
	res = append(res, idx)
	for nbr > 0 {
		idx = nbr % baseLen
		res = append(res, idx)
		nbr /= baseLen
	}
	if isNegative {
		z01.PrintRune('-')
	}
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(rune(base[res[i]]))
	}
}
