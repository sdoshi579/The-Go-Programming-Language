package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	unique := make(map[string]bool)

	for input.Scan() {
		unique[input.Text()] = true
	}
	fmt.Println(unique)
}
