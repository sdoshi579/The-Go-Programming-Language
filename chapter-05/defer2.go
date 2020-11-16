// I'm not sure how this work if the painstack catches
// the panic then function should continue to run after that which is not happening
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer painstack()
	f(3)
}

func painstack() {
	var buf [4096]byte
	// stack th e
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
