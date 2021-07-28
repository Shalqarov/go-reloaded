package student

func ListPushFront(l *List, data int) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
	} else {
		temp := &NodeL{Data: data}
		temp.Next = l.Head
		l.Head = temp
	}
}
