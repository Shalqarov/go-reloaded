package student

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data_ref {
		if l.Head.Data == data_ref && l.Head != nil {
			l.Head = l.Head.Next
		}
		if l.Head == nil {
			return
		}
	}
	cur := l.Head.Next
	prev := l.Head
	for cur != nil {
		if cur.Data == data_ref {
			for cur != nil && cur.Data == data_ref {
				cur = cur.Next
			}
			prev.Next = cur
			if cur == nil {
				return
			}
		}
		cur = cur.Next
		prev = prev.Next
	}
}
