package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
	"github.com/01-edu/z01.PrintRune"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
	}else{
		args:=os.Args
		vowels:=[]rune{}
		for i:=0;i<len(args);i++{
			vowels = []rune{}
			for j:=0;j<len(args[i]);j++{
				if args[i][j] == 
			}
		}
	}
}