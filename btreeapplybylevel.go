package student

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
