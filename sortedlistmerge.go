package student

func NodeSize(n *NodeI) int {
	cnt := 0
	iterator := n
	for iterator != nil {
		cnt++
		iterator = iterator.Next
	}
	return cnt
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 == nil {
		return nil
	}
	var result = &NodeI{}
	if n1 != nil {
		result = &NodeI{Data: n1.Data}
		n1 = n1.Next
	} else {
		result = &NodeI{Data: n2.Data}
		n2 = n2.Next
	}
	for {
		if n1 != nil {
			SortListInsert(result, n1.Data)
			n1 = n1.Next
		}
		if n2 != nil {
			SortListInsert(result, n2.Data)
			n2 = n2.Next
		}
		if n1 == nil && n2 == nil {
			break
		}
	}
	return result
}
