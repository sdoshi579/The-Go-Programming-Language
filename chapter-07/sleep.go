//go run chapter-07/sleep.go -period 3s
package main

import (
	"flag"
	"fmt"
	"time"
)

// Create flag name period with type duration
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
