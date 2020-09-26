// add comma before last three characters of any strings
package main

import (
	"bufio"
	"bytes"
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
	var buf bytes.Buffer
	length := len(s)
	if length <= 3 {
		buf.WriteString(s)
		return buf.String()
	}

	buf.WriteString(s[:length-3])
	buf.WriteString(",")
	buf.WriteString(s[length-3:])
	return buf.String()
}
