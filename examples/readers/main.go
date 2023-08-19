package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Reader")

	bytes := make([]byte, 8)
	for {
		n, err := reader.Read(bytes)
		fmt.Printf("n = %v err = %v bytes = %v\n", n, err, bytes)
		fmt.Printf("bytes[:n] = %q\n", bytes[:n])

		if err == io.EOF {
			break
		}
	}
}
