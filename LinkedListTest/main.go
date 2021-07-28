package main

import (
	"fmt"
	"student"
)

func main() {
	// var link *student.NodeI

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 1)
	// link = listPushBack(link, 2)
	// link = listPushBack(link, 3)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 5)
	// link = listPushBack(link, 24)
	// link = listPushBack(link, 25)
	// link = listPushBack(link, 54)
	// PrintList(link)

	// // link = student.SortListInsert(link, 0)
	// link = student.SortListInsert(link, 33)
	// PrintList(link)

	// var link *student.NodeI
	// var link2 *student.NodeI

	// link = listPushBack(link, 2)
	// link = listPushBack(link, 7)
	// link = listPushBack(link, 39)
	// link = listPushBack(link, 92)
	// link = listPushBack(link, 97)
	// link = listPushBack(link, 93)
	// link = listPushBack(link, 91)
	// link = listPushBack(link, 28)
	// link = listPushBack(link, 64)

	// link2 = listPushBack(link2, 2)
	// link2 = listPushBack(link2, 2)
	// link2 = listPushBack(link2, 4)
	// link2 = listPushBack(link2, 9)
	// link2 = listPushBack(link2, 12)
	// link2 = listPushBack(link2, 12)
	// link2 = listPushBack(link2, 19)
	// link2 = listPushBack(link2, 20)
	// // PrintList(link)
	// // PrintList(link2)
	// PrintList(student.SortedListMerge(link2, link))

	link := &student.List{}
	link2 := &student.List{}

	fmt.Println("----normal state----")
	//student.ListPushBack(link2, 1)
	PrintList(link2)
	student.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "Hello")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "There")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "How")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "are")
	student.ListPushBack(link, "you")
	student.ListPushBack(link, 1)
	PrintList(link)

	student.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)

}

// func PrintList(l *student.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

func PrintList(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

// type node struct {
// 	data int
// 	next *node
// }

// type linkedList struct {
// 	head   *node
// 	length int
// }

// func (l *linkedList) prepend(n *node) {
// 	second := l.head
// 	l.head = n
// 	l.head.next = second
// 	l.length++
// }

// func (l linkedList) printListData() {
// 	toPrint := l.head
// 	for l.length != 0 {
// 		fmt.Printf("%d ", toPrint.data)
// 		toPrint = toPrint.next
// 		l.length--
// 	}
// 	fmt.Println()
// }

// func (l *linkedList) deleteWithValue(value int) {
// 	if l.length == 0 {
// 		return
// 	}
// 	if l.head.data == value {
// 		l.head = l.head.next
// 		l.length--
// 		return
// 	}
// 	previousToDelete := l.head
// 	for previousToDelete.next.data != value {
// 		if previousToDelete.next.next == nil {
// 			return
// 		}
// 		previousToDelete = previousToDelete.next
// 	}
// 	previousToDelete.next = previousToDelete.next.next
// 	l.length--
// }

// func main() {
// 	mylist := linkedList{}
// 	node1 := &node{data: 48}
// 	node2 := &node{data: 18}
// 	node3 := &node{data: 16}
// 	node4 := &node{data: 16}
// 	node5 := &node{data: 11}
// 	node6 := &node{data: 7}
// 	node7 := &node{data: 2}
// 	mylist.prepend(node1)
// 	mylist.prepend(node2)
// 	mylist.prepend(node3)
// 	mylist.prepend(node4)
// 	mylist.prepend(node5)
// 	mylist.prepend(node6)
// 	mylist.prepend(node7)
// 	mylist.printListData()
// 	mylist.deleteWithValue(48)
// 	mylist.printListData()

// }
