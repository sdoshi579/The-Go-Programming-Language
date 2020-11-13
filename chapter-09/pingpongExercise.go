package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	var i int
	q := make(chan int)
	start := time.Now()
	go func() {
		q <- 1
		for {
			i++
			q <- <-q
		}
	}()
	go func() {
		for {
			q <- <-q
		}
	}()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println(float64(i)/float64(time.Since(start))*1e9, "communications per second")
}
