// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)

}

func reverse(s *[6]int) {
	n := len(s) - 1
	for i := range s {
		if i > n/2 {
			break
		}
		s[i], s[n-i] = s[n-i], s[i]
	}
}
