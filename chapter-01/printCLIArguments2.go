// prints command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	// for first initialization, condition, post loop computation
	for _, i := range os.Args[1:] {
		// os.Args[0] contains the command name itself
		s += sep + i
		sep = " "
	}
	fmt.Println(s)
}
