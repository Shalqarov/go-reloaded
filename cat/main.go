package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	argsSize := len(args)
	if argsSize == 0 {
		fmt.Println("entr")
		//вывод введенной строки
		data := make([]byte, 0)
		_, err := os.Stdin.Read(data)
		if err == nil {
			datalen := len(data)
			for i := 0; i < datalen-1; i++ {
				z01.PrintRune(rune(data[i]))
			}
		} else {
			os.Exit(1)
		}

	} else {
		for i := 0; i < argsSize; i++ {
			msg := []rune{}
			file, err := os.Open(args[i])
			if err != nil {
				msg := []rune{}

				for i := 0; i < len(msg); i++ {

				}
				os.Exit(1)
			}
			f, err := file.Stat()
			arr := make([]byte, f.Size())
			file.Read(arr)
			for i := 0; i < len(arr); i++ {
				z01.PrintRune(rune(arr[i]))
			}
			z01.PrintRune('\n')
		}
	}
}
