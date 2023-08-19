package main

import (
	"fmt"

	"example.com/linkedlist"
)

func main() {
	l := linkedlist.LinkedList[int]{}
	l.Append(5).Append(7).Append(10)
	fmt.Println(l)
}
