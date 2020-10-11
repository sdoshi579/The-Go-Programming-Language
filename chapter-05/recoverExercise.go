package main

import (
	"fmt"
)

func main() {
	fmt.Println(f())
}

func f() (t string) {
	defer func() {
		if p := recover(); p != nil {
			t = "recover"
		}
	}()
	panic("test")
}
