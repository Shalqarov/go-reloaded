package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	argsSize := len(args)
	if argsSize == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for i := 0; i < argsSize; i++ {
			str, err := ioutil.ReadFile(args[i])
			if err != nil {
				printErr("ERROR: " + err.Error() + "\n")
				os.Exit(0)
			}
			printErr(string(str))
		}
	}
}

func printErr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}
