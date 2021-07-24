package student

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if root.Data == elem {
			return root
		} else if root.Data < elem {
			return BTreeSearchItem(root.Right, elem)
		} else {
			return BTreeSearchItem(root.Left, elem)
		}
	}
	return nil
}
