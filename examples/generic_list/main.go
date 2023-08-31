package main

import (
	"fmt"

	"example.com/linkedlist"
)

func main() {
	l := linkedlist.LinkedList[int]{}
	l.Append(5)
	l.Append(7)
	l.Append(10)
	l.RemoveLast()
	fmt.Println(l)
}
