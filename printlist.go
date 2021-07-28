package student

import "fmt"

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}
