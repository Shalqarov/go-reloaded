package student

func AtoiBase(s string, base string) int {
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
