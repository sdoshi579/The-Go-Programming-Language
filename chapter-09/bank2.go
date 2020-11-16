package main

import (
	"fmt"
)

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func main() {
	done := make(chan struct{})

	go func() {
		deposit(500)
		fmt.Printf("balance: %d\n", Balance())
		done <- struct{}{}
	}()

	go func() {
		deposit(1000)
		fmt.Printf("balance: %d\n", Balance())
		done <- struct{}{}
	}()

	<-done
	<-done
}
