// Fetchall fetches URLs in parallel and reports their times and sizes.
// The content we get each time is same and I can't see any caching mechanism as running continously is still taking same time
// With new websites I tried the command go run chapter-01/fetchAllExercise.go https://golang.org http://gopl.io https://godoc.org https://alexa.com https://ahrefs.com https://similarweb.com https://backlinko.com https://semrush.com
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	//Discard is a io writer type but does not write
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("======================\n %.2fs\n  %s\n  %s\n =============================", secs, b, url)
}
