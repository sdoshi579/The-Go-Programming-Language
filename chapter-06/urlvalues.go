package main

import "fmt"

// Values is map of string and string array
type Values map[string][]string

func main() {
	m := Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3") // causes panic as v is nil
}

// Get returns the value of given key
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add appends the value to given key
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
