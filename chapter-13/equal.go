package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type // Pointer can be same so need type too for example: x and x[0] has same ptr
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())

		if xptr == yptr {
			return true
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Int:
		return x.Int() == y.Int()
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return x.Pointer() == y.Pointer()
	case reflect.Interface, reflect.Ptr:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func main() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal("foo", "bar"))
	fmt.Println(Equal([]string(nil), []string{}))

}
