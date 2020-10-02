package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("a", "b")
	addEdge("c", "d")
	addEdge("a", "d")
	addEdge("d", "a")
	fmt.Println(fmt.Sprintf("Edge between a and b: %v", hasEdge("a", "b")))
	fmt.Println(fmt.Sprintf("Edge between c and d: %v", hasEdge("c", "d")))
	fmt.Println(fmt.Sprintf("Edge between d and a: %v", hasEdge("d", "a")))
	fmt.Println(fmt.Sprintf("Edge between a and d: %v", hasEdge("a", "d")))
	fmt.Println(fmt.Sprintf("Edge between x and b: %v", hasEdge("x", "b")))
	fmt.Println(fmt.Sprintf("Edge between x and d: %v", hasEdge("x", "d")))
	fmt.Println(fmt.Sprintf("Edge between d and x: %v", hasEdge("d", "x")))
}
