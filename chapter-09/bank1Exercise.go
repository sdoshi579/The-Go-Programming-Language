package main

import (
	"fmt"
)

type withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdrawals = make(chan withdrawal)

func deposit(amount int) {
	deposits <- amount
}

func balance() int {
	return <-balances
}

func withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- withdrawal{amount, ch}
	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdrawStruct := <-withdrawals:
			if balance >= withdrawStruct.amount {
				balance -= withdrawStruct.amount
				withdrawStruct.success <- true
			} else {
				withdrawStruct.success <- false
			}
		}
	}
}

func main() {
	go teller()
	done := make(chan struct{})

	go func() {
		deposit(500)
		fmt.Printf("balance: %d\n", balance())
		done <- struct{}{}
	}()

	go func() {
		deposit(1000)
		fmt.Printf("balance: %d\n", balance())
		done <- struct{}{}
	}()

	go func() {
		succ := withdraw(500)
		fmt.Printf("amount deducted status: %v current balance: %d\n", succ, balance())
		done <- struct{}{}
	}()

	go func() {
		succ := withdraw(1000)
		fmt.Printf("amount deducted status: %v current balance: %d\n", succ, balance())
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done
	<-done
}
