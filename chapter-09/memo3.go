package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Memo struct {
	f     Func
	cache map[string]result
	mu    sync.Mutex
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func new(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

func (memo *Memo) get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	if resp, err := http.Get(url); err != nil {
		return nil, err
	} else {
		return ioutil.ReadAll(resp.Body)
	}
}

func main() {
	m := new(httpGetBody)
	var n sync.WaitGroup
	for _, url := range os.Args[1:] {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.get(url)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
