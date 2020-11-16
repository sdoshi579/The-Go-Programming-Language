package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'e', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.Type().String()
	default:
		return v.Type().String()
	}
}

func main() {
	fmt.Println(any(1))
	fmt.Println(any(true))
	fmt.Println(any([]string{"hello", "world"}))
	fmt.Println(any(map[int]string{1: "hello"}))
	fmt.Println(any(1.12))
	fmt.Println(any(-1))
	fmt.Println(any("string test"))
}
