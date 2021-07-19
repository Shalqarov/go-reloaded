package student

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	strLen := len(a)
	for i := 0; i < strLen; i++ {
		for j := 0; j < strLen-1; j++ {
			if f(a[j], a[i]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
