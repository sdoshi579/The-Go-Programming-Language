package main

import (
	"fmt"
	"bytes"
	"io"
	"strings"
)

type limitReader struct {
	reader io.Reader
	limit int
	count int
} 

func main() {
	l := limitReader{
		reader: strings.NewReader("hello world"),
		limit: 5,
	}
	scan := bytes.Buffer{}
	scan.ReadFrom(&l)
	fmt.Println(scan.String())
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	n, err = l.reader.Read(p[:l.limit])
	l.count  = n
	if l.count >= l.limit {
		err = io.EOF
	}
	return
}