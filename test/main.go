package main

import (
	"fmt"
	"strings"
	"student"
)

func main() {
	fmt.Println(student.Split("HelloHAHowHAare you?HAHAHAA", "HA"))
	fmt.Println(strings.Split("HelloHAHowHAare you?HAHAHAA", "HA"))
}
