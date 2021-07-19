package main

import (
	"fmt"
	"student"
)

func main() {

	result := []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	student.AdvancedSortWordArr(result, student.Compare)
	fmt.Println(result)
}
