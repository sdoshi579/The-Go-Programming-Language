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
	requests chan request
}

type request struct {
	key      string
	response chan<- result
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
	memo := &Memo{
		requests: make(chan request),
	}
	go memo.server(f)
	return memo
}

func (memo *Memo) get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{
				ready: make(chan struct{}),
			}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
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
