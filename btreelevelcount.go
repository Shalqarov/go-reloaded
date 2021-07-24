package student

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
