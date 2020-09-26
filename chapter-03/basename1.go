package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	forwardSlash := strings.LastIndex(s, "/")
	s = s[forwardSlash+1:]
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		s = s[:dot]
	}
	return s
}
