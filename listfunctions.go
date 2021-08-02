package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

type NodeI struct {
	Data int
	Next *NodeI
}

func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
	} else {
		temp := l.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &NodeL{Data: data}
	}
}

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

func ListClear(l *List) {
	l.Head = nil
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head.Next != nil {
		l.Head = l.Head.Next
	}
	return l.Head.Data
}

func ListPushFront(l *List, data int) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
	} else {
		temp := &NodeL{Data: data}
		temp.Next = l.Head
		l.Head = temp
	}
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}
	if l.Head == nil {
		return
	}
	if l.Head.Data == data_ref {
		if l.Head.Data == data_ref && l.Head != nil {
			l.Head = l.Head.Next
		}
		if l.Head == nil {
			l.Tail = nil
			return
		}
	}
	cur := l.Head.Next
	prev := l.Head
	for cur != nil {
		if prev.Data == data_ref {
			if l.Head.Data == data_ref && l.Head != nil {
				l.Head = l.Head.Next
			}
			if l.Head == nil {
				l.Tail = nil
				return
			}
		}
		if cur.Data == data_ref {
			for cur != nil && cur.Data == data_ref {
				cur = cur.Next
			}
			prev.Next = cur
			if cur == nil {
				l.Tail = prev
				return
			}
		}
		cur = cur.Next
		prev = prev.Next
	}
}

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	prev := &NodeL{Data: l.Head.Data, Next: nil}
	ListLength := ListSize(l)
	for i := 0; i < ListLength; i++ {
		current := l.Head
		for current.Next != nil && current.Next != prev {
			current.Data, current.Next.Data = current.Next.Data, current.Data
			current = current.Next
		}
		prev = current
	}
}

func ListSize(l *List) int {
	cnt := 0
	for l.Head != nil {
		cnt++
		l.Head = l.Head.Next
	}
	return cnt
}

func FindTail(n *NodeI) *NodeI {
	if n == nil {
		return nil
	}
	iterator := n
	for {
		iterator = iterator.Next
		if iterator.Next == nil {
			return iterator
		}
	}
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 == nil {
		return nil
	}
	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}
	tail := FindTail(n1)
	if tail != nil {
		tail.Next = n2
	}
	//sorting...
	current := n1
	temp := n1
	for current.Next != nil {
		temp = current.Next
		for temp != nil {
			if temp.Data < current.Data {
				t := temp.Data
				temp.Data = current.Data
				current.Data = t
			}
			temp = temp.Next
		}
		current = current.Next
	}
	return n1
}

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
