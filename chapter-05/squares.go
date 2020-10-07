package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Printf("%#v\n", f)
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}
}
