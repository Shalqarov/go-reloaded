package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	largs := len(args)
	filenames := []string{}
	bytes := 0
	plus := false

	for i := 0; i < largs; i++ {
		if args[i] == "-c" && i+1 != largs {
			if isNum(args[i+1]) {
				if len(args[i+1]) > 19 && args[i+1] != "18446744073709551615" && args[i+1][0] != '-' {
					fmt.Printf("tail: invalid number of bytes: ‘%s’: Value too large for defined data type\n", args[i+1])
					return
				}
				if isPlus(args[i+1]) {
					plus = true
				}
				bytes = int(Atoi(args[i+1]))

				if bytes < 0 {
					bytes *= -1
				}
				i++
				continue
			} else {
				fmt.Printf("tail: invalid number of bytes: '%s'\n", args[i+1])
				return
			}
		} else if args[i] == "-c" && i+1 == largs {
			fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
			return
		}
		filenames = append(filenames, args[i])

	}

	notfirst := false

	for _, ch := range filenames {
		file, err := os.Open(ch)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", ch)
			continue
		}

		fstat, _ := file.Stat()
		size := fstat.Size()

		res := make([]byte, size)
		file.Read(res)
		start := 0

		if plus {
			start = len(res) - (len(res) - bytes) - 1
		} else {
			start = len(res) - bytes
		}

		if start > len(res) {
			return
		}

		if start < 0 {
			start = 0
		}
		if len(filenames) > 1 && !notfirst {
			fmt.Printf("==> %s <==\n%s", ch, res[start:])
		} else if len(filenames) > 1 && notfirst {
			fmt.Printf("\n==> %s <==\n%s", ch, res[start:])
		} else {
			fmt.Printf("%s", res[start:])
		}
		notfirst = true
	}

}

func isNum(s string) bool {
	if s == "" {
		return false
	}

	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return true
}

func isPlus(s string) bool {
	if isNum(s) {
		if s[0] == '+' {
			return true
		}
	}
	return false
}

func Atoi(s string) int64 {
	var res int64
	s1 := s
	if len(s) == 0 {
		return 0
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	for _, ch := range s {
		if ch >= 48 && ch <= 57 {
			ch = ch - 48
			res = res*10 + int64(ch)
			if res < 0 {
				break
			}
		} else {
			return 0
		}
	}
	if s1[0] == '-' {
		res = -res
	}

	if (s1[0] == '+' || (s1[0] >= '0' && s1[0] <= '9')) && res < 0 {
		return 9223372036854775807
	} else if s1[0] == '-' && res > 0 {
		return -9223372036854775808
	}

	return res
}
