package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:]
	if len(input) == 1 {
		fmt.Printf("%x\n", sha256.Sum256([]byte(input[0])))
		os.Exit(1)
	}

	if input[1] == "384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(input[0])))
	} else if input[1] == "512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(input[0])))
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(input[0])))
	}
}
