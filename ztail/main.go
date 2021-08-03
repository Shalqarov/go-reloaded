package main

import (
	"fmt"
	"os"
	"student"
)

func main() {
	args := os.Args[1:]
	argsSize := len(args)
	if argsSize == 0 {
		os.Exit(0)
	}
	if argsSize < 1 {
		os.Exit(0)
	} else {
		bytes := 0    // кол-во байт
		plus := false //проверка на +
		filenames := []string{}
		for i := 0; i < argsSize; i++ {
			if args[i] == "-c" && i+1 != argsSize {
				if isNum(args[i+1]) {
					if isPlus(args[i+1]) {
						plus = true
					} else {
						plus = false
					}
					bytes = student.Atoi(args[i+1])
					if bytes < 0 {
						bytes *= -1
					}
					i++
					continue
				} else {
					fmt.Printf("tail: invalid number of bytes: '%s'\n", args[i+1])
					os.Exit(0)
				}
			} else if args[i] == "-c" && i+1 == argsSize {
				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information\n")
				os.Exit(0)
			}
			filenames = append(filenames, args[i])
		}
		//fmt.Println(filenames)	// проверяю правильно ли работает переменная (нужно удалить строку)
		if len(filenames) == 1 {
			file, err := os.Open(filenames[0])
			if err != nil {
				fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", filenames[0])
				os.Exit(0)
			}
			fs, _ := file.Stat()
			fileSize := fs.Size()
			res := make([]byte, fileSize)
			start := 0
			if plus {
				start = len(res) - (len(res) - bytes)
			} else {
				start = len(res) - bytes
			}
			if start < 0 {
				start = 0
			} else if start >= len(res) {
				start = 0
			}
			file.Read(res)
			if plus {
				start = start - 1
				if start < 0 {
					start = 0
				}
			}
			fmt.Printf("%s", res[start:])

		} else {
			prevErr := false
			firstErr := false
			for i, s := range filenames {
				file, err := os.Open(s)
				if err != nil {
					prevErr = true
					if i == 0 {
						firstErr = true
					}
					fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", s)
					continue
				}
				fs, _ := file.Stat()
				fileSize := fs.Size()
				res := make([]byte, fileSize)
				start := 0
				if plus {
					start = len(res) - (len(res) - bytes)
				} else {
					start = len(res) - bytes
				}
				if start < 0 {
					start = 0
				}
				file.Read(res)
				if plus {
					start = start - 1
					if start < 0 {
						start = 0
					}
				}
				if i > 0 {
					if prevErr == true && !firstErr {
						fmt.Printf("\n")
						prevErr = false
					}
					fmt.Printf("==> %s <==\n%s", s, res[start:])
				} else {
					fmt.Printf("==> %s <==\n%s", s, res[start:])
				}
				firstErr = false
			}
		}
	}
}

func isPlus(s string) bool {
	if !isNum(s) {
		return false
	}
	if s[0] == '+' {
		return true
	}
	return false
}

func isNum(s string) bool {
	if s == "" {
		return false
	}
	lS := len(s)
	start := 0
	if s[0] == '-' {
		start = 1
	}
	if s[0] == '+' {
		start = 1
	}
	for i := start; i < lS; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
