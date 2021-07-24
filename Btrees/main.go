package main

import (
	"fmt"
	"student"
)

// func BtreeTransplant(root, node, rplc *TreeNode) *TreeNode {

// }

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "2")
	student.BTreeInsertData(root, "3")
	student.BTreeInsertData(root, "8")
	student.BTreeApplyByLevel(root, fmt.Println)

}
