package main

import (
	"io"
	"os"
	"strings"
)

func cipherLetter(b byte) byte {
	baseLowerByte := byte('a')
	baseUpperByte := byte('A')
	resultByte := byte(0)
	if b >= 'a' && b <= 'z' {
		resultByte = ((b-baseLowerByte)+13+26)%26 + baseLowerByte
	} else if b >= 'A' && b <= 'Z' {
		resultByte = ((b-baseUpperByte)+13+26)%26 + baseUpperByte
	} else {
		resultByte = b
	}

	return resultByte
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rot13.r.Read(bytes)

	if err == io.EOF {
		return 0, io.EOF
	}

	for i := 0; i < n; i++ {
		bytes[i] = cipherLetter(bytes[i])
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
