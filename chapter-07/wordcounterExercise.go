package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type wordCounter int

func main() {
	var w wordCounter
	w.Write([]byte("This is a word counter exercise"))
	fmt.Println(w)

	w = 0
	fmt.Fprintf(&w, "This is a word counter exercise\n Exercise for byte counter")
	fmt.Println(w)
}

func (w *wordCounter) Write(p []byte) (int, error) {
	bytescanner := bufio.NewScanner(bytes.NewReader(p))
	bytescanner.Split(bufio.ScanWords)
	for bytescanner.Scan() {
		*w++
	}
	return len(p), nil
}
