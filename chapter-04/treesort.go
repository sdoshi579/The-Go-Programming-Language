package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	a := make([]int, 5)
	i := 0
	for input.Scan() {
		x, _ := strconv.Atoi(input.Text())
		a[i] = x
		i++
	}
	fmt.Println(a)
	sort(a)
	fmt.Println(a)
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
