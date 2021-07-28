package student

type NodeL struct {
	//Data interface{}
	Data int
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

func ListPushBack(l *List, data int) {
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

func ListPushBackN(l *NodeI, data int) {
	if l == nil {
		l = &NodeI{Data: data}
	} else {
		temp := l
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &NodeI{Data: data}
	}
}
