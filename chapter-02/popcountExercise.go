// Popcount exercise 2.3
package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func main() {
	x := os.Args[1]
	j, _ := strconv.ParseUint(x, 10, 64)
	fmt.Println(popcount(j))
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popcount(x uint64) int {
	var t byte
	for i := 0; i < 8; i++ {
		t += pc[byte(x>>(i*8))]
	}
	return int(t)
}
