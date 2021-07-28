package student

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil {
		return newNode
	}
	if l.Data > data_ref {
		newNode.Next = l
		return newNode
	}
	prev := l
	cur := l.Next
	for cur != nil {
		if cur.Data >= data_ref {
			newNode.Next = cur
			break
		}
		prev = cur
		cur = cur.Next
	}
	prev.Next = newNode
	return l
}
