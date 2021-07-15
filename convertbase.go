package student

func atoibase(s string, base string) int {
	if s == "" || len(base) < 2 {
		return 0
	}
	indexes := make(map[rune]int, 0)
	for i, symbol := range base {
		if _, isHave := indexes[symbol]; isHave || symbol == '-' || symbol == '+' {
			return 0
		}
		indexes[symbol] = i
	}
	baseLen := len(base)
	num := 0
	for _, d := range s {
		num = num*baseLen + indexes[d]
	}
	return num
}

func printnbr(nbr int, base string) string {
	baseLen := len(base)
	idx := nbr % baseLen
	nbr /= baseLen
	res := []int{}
	res = append(res, idx)
	for nbr > 0 {
		idx = nbr % baseLen
		res = append(res, idx)
		nbr /= baseLen
	}
	temp := 0
	for i, j := 0, len(res)-1; i < j; i++ {
		temp = res[i]
		res[i] = res[j]
		res[j] = temp
		j--
	}
	strRes := ""
	for i := 0; i < len(res); i++ {
		strRes += string(base[res[i]])
	}
	return strRes
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := atoibase(nbr, baseFrom)
	return printnbr(decimal, baseTo)
}
