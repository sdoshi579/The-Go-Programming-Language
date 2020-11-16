// add comma before last three characters of any strings
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}

func comma(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	}

	return s[:length-3] + "," + s[length-3:]
}
