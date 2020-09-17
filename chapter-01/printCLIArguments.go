// prints command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		// os.Args[0] contains the command name itself
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
