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
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
