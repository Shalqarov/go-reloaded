package main

import "os"

func main() {
	args := os.Args[1:]
	argsSize := len(args)
	if argsSize < 3 {
		os.Exit(0)
	} else {
		Carg := false
		for i := 0; i < argsSize; i++ {
			if args[i] == "-c" {
				Carg = true
				if i+1 == argsSize {
					os.Exit(0)
				}
			}
			if !Carg {
				os.Exit(0)
			}
		}
	}
}
