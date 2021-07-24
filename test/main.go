package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println(N)
	var nodes []Node = make([]Node, N)
	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		nodes[i].Value = val
		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
		for _, node := range nodes {
			printNode(&node)
		}
	}
}

func printNode(n *Node) {
	fmt.Print("Value: ", n.Value)
	if n.Left != nil {
		fmt.Print(" Left: ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Print(" Right: ", n.Right.Value)
	}

	fmt.Println()
}

// func main() {
// 	test := &Node{
// 		Value: 123,
// 	}
// 	right := &Node{
// 		Value: 333,
// 	}
// 	test.Right = right
// 	fmt.Println((test))
// 	printNode(test)
// }
