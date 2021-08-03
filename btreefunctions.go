package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

// ========>    ApplyByLevel    <==========
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	lvls := BTreeLevelCount(root)
	for i := 1; i <= lvls; i++ {
		CurrentLevel(root, i, f)
	}
}

func CurrentLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		CurrentLevel(root.Left, level-1, f)
		CurrentLevel(root.Right, level-1, f)
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyPostorder(root.Left, f)
		BTreeApplyPostorder(root.Right, f)
		f(root.Data)
	}
}

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right, f)

	}
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
		if root.Right != nil {
			root.Right.Parent = root
		}
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			temp := BTreeMin(root.Right)
			root.Data = temp.Data
			root.Right = BTreeDeleteNode(root.Right, temp)
			if root.Right != nil {
				root.Right.Parent = root
			}
		}
	}
	return root
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil || root.Right == nil || root.Left == nil {
		return true
	}
	checker := false
	if root.Left.Data < root.Data && root.Right.Data > root.Data {
		checker = BTreeIsBinary(root.Right)
		if checker == true {
			checker = BTreeIsBinary(root.Left)
		}
	}
	return checker
}

func BTreeLevelCount(root *TreeNode) int {
	cnt := 0
	if root == nil {
		return cnt
	}
	cnt++
	cntLeft := cnt + BTreeLevelCount(root.Left)
	cntRight := cnt + BTreeLevelCount(root.Right)
	if cntRight >= cntLeft {
		return cntRight
	} else {
		return cntLeft
	}
}

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

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

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	branch := BTreeSearchItem(root, node.Data)
	if branch == nil || branch.Parent == nil {
		return rplc
	}
	parent := branch.Parent
	if node.Data < parent.Data {
		parent.Left = rplc
	} else {
		parent.Right = rplc
	}
	rplc.Parent = parent
	return root
}
