package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	rotate(a, 1)
	fmt.Println(a)
}

func rotate(a []int, order int) {
	tail := make([]int, order)
	copy(tail, a[0:order])
	copy(a, a[order:])
	for i, n := range tail {
		a[len(a)-order+i] = n
	}
}
