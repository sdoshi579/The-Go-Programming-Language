// prints number of differnt bits between hash of x and X
package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var count int

	for i := range c1 {
		count += bits.OnesCount(uint(c1[i] ^ c2[i]))
	}
	fmt.Println(count)
	// fmt.Println(bits.OnesCount(uint(c1[0] ^ c2[0])))
	// fmt.Printf("%08b\t %v\t %x\n", c1[0], c1[0], c1[0])
	// fmt.Printf("%08b\t %v\t %x\n", c2[0], c2[0], c2[0])
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
