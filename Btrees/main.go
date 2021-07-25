package main

import (
	"fmt"
	"student"
)

// func BtreeTransplant(root, node, rplc *TreeNode) *TreeNode {

// }

func main() {
	root := &student.TreeNode{Data: "33"}
	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "20")
	student.BTreeInsertData(root, "31")
	student.BTreeInsertData(root, "52")
	student.BTreeInsertData(root, "13")
	student.BTreeInsertData(root, "11")
	node := student.BTreeSearchItem(root, "20")
	fmt.Println("Before delete:")
	student.BtreePrint(root)
	// student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BtreePrint(root)

	// student.BTreeApplyInorder(root, fmt.Println)
}
