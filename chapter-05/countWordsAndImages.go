package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		words, images, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			continue
		}
		fmt.Printf("%d \t %d \n", words, images)
	}
}

func findLinks(url string) (int, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return 0, 0, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	t, t1 := visit1(doc)
	return t, t1, nil
}

func visit1(n *html.Node) (words, img int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		img = 1
	}
	if n.Type == html.TextNode {
		words = wordCount(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := visit1(c)
		img += i
		words += w
	}
	return words, img
}

func wordCount(text string) int {
	words := strings.Split(text, " ")
	return len(words)
}
