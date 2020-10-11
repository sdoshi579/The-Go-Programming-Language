package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	// We deferred the function returned from trace
	// If we use defer trace("bigSlowOperation") then we deferred the function call and vever invoked the return function
	defer trace("bigSlowOperation")()

	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
