// go run chapter-01/fetch.go https://golang.org | go run chapter-05/findlinks.go
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/*
//!+html
package html
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32
const (
	ErrorNode NodeType = iota
	TextNode <textarea>
	DocumentNode
	ElementNode <a><br>
	CommentNode <!--- !--->
	DoctypeNode <!Doctype>
)
type Attribute struct {
	Key, Val string  //<img color=#fffff> key = color and val = #fffff
}
func Parse(r io.Reader) (*Node, error)
*/
