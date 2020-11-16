package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n ", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}

func main() {
	print(time.Hour)
	print(http.NewServeMux())
}
