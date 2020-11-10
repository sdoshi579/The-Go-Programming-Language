package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
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
