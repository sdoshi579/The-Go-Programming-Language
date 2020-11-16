package main

import "fmt"

type byteCounter int

func main() {
	var b byteCounter
	b.Write([]byte("hello world"))
	fmt.Println(b)

	b = 0
	fmt.Fprintf(&b, "world hello")
	fmt.Println(b)
}

func (b *byteCounter) Write(p []byte) (int, error) {
	*b += byteCounter(len(p))
	return len(p), nil
}
