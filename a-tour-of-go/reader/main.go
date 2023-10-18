package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

// Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
	b := make([]byte, 1)
	cnt := 0
	for cnt <= 100 {
		n, err := MyReader{}.Read(b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
		cnt++
	}
}
