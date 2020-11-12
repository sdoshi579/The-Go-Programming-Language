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
	cache map[string]*entry
	mu    sync.Mutex
}

type entry struct {
	res   result
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func new(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

func (memo *Memo) get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{
			ready: make(chan struct{}),
		}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
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
