package student

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
	tail := FindTail(n1)
	tail.Next = n2
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
