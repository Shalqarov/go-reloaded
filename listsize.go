package student

func ListSize(l *List) int {
	cnt := 0
	for l.Head != nil {
		cnt++
		l.Head = l.Head.Next
	}
	return cnt
}
