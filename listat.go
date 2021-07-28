package student

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	iterator := l
	for i := 0; iterator != nil; iterator = iterator.Next {
		if i == pos {
			return iterator
		}
		i++
	}
	return nil
}
