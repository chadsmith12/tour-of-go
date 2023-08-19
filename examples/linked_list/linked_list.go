package linkedlist

import "fmt"

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (list *LinkedList[T]) Append(value T) *LinkedList[T] {
	node := &Node[T]{Value: value}
	if list.head == nil {
		list.head = node
		return list
	}
	currentNode := list.head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = node
	return list
}

func (list LinkedList[T]) String() string {
	if list.head == nil {
		return fmt.Sprintf("%v", list.head)
	}

	currentNode := list.head
	stringified := fmt.Sprintf("%v->", currentNode.Value)
	for currentNode.Next != nil {
		stringified += fmt.Sprintf("%v->", currentNode.Next.Value)
		currentNode = currentNode.Next
	}
	stringified += fmt.Sprintf("%v\n", currentNode.Next)

	return stringified
}
