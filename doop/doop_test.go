package main //run via "go test"

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

//**************run via: "go test -v" command*****************
func TestMain(t *testing.T) {

	prevArgs := os.Args
	defer func() { os.Args = prevArgs }()

	testCases := []struct {
		v1       string
		operator string
		v2       string
		result   string
	}{
		{"1", "+", "1", "2\n"},
		{"hello", "+", "1", "0\n"},
		{"1", "p", "1", "0\n"},
		{"1", "/", "0", "No division by 0\n"},
		{"1", "%", "0", "No modulo by 0\n"},
		{"9223372036854775807", "+", "1", "0\n"},
		{"-9223372036854775809", "-", "3", "0\n"},
		{"9223372036854775807", "*", "3", "0\n"},
		{"861", "+", "870", "1731\n"},
		{"861", "-", "870", "-9\n"},
		{"861", "*", "870", "749070\n"},
		{"861", "%", "870", "861\n"},
		{"-9223372036854775808", "-", "1", "0\n"},
		{"-9223372036854775808", "*", "1", "-9223372036854775808\n"},
		{"-9223372036854775808", "*", "-1", "0\n"},
		{"-9223372036854775808", "/", "1", "-9223372036854775808\n"},
		{"-9223372036854775808", "/", "-1", "0\n"},
		{"-9223372036854775808", "+", "-1", "0\n"},
		{"-9223372036854775808", "-", "-1", "-9223372036854775807\n"},
		{"9223372036854775807", "-", "-1", "0\n"},
		{"9223372036854775807", "/", "9223372036854775807", "1\n"},
		{"9223372036854775807", "*", "9223372036854775807", "0\n"},
		{"9223372036854775807", "+", "9223372036854775807", "0\n"},
		{"1", "+", "+", "0\n"},
		{"9223372036854775656", "*", "3", "0\n"},
	}
	errorcount := 0
	for _, test := range testCases {
		os.Args = append(os.Args, []string{test.v1, test.operator, test.v2}...)

		got := captureOutput(main)
		fmt.Printf("got: %s", got)
		// main()
		if test.result != string(got) {
			t.Errorf("wait %v  %v  %v  = %v but got %v", test.v1, test.operator, test.v2, test.result, got)
			errorcount++
		}
	}
	if errorcount == 0 {
		fmt.Printf("Doop - passed\n")
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// func captureOutput(f func()) string {
// 	var buf bytes.Buffer
// 	log.SetOutput(&buf)
// 	f()
// 	log.SetOutput(os.Stdout)
// 	a := buf.String()
// 	fmt.Println(a)
// 	return buf.String()
// }
