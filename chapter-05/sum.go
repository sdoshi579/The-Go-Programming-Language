package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(2))
	fmt.Println(sum(1, 2, 3, 4))
}

func sum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
