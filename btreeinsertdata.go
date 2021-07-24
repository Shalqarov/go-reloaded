package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if root.Data < data {
		//move right
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	} else if root.Data > data {
		//move left
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	}
	return root
}
