// Check if both the strings are anagrams
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s1, s2 := os.Args[1], os.Args[2]

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if len(s1) != len(s2) {
		fmt.Println("Not anagrams")
		os.Exit(1)
	}

	//rune1 := []rune(s1)
	rune2 := []rune(s2)

	for _, r := range rune2 {
		if strings.ContainsRune(s1, r) {
			continue
		}
		fmt.Println("Not anagrams")
		os.Exit(1)
	}

	fmt.Println("Both strings are anagrams")
}
