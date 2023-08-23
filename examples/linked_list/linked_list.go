package linkedlist

import (
	"fmt"
	"errors"
)

type Node[T any] struct {
	Next  *Node[T]
	Prev *Node[T]
	Value T
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	length int
}

func (list *LinkedList[T]) Append(value T) *LinkedList[T] {
	node := &Node[T]{Value: value}
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length = 1
		return list
	}
	list.length += 1
	list.tail.Next = node
	node.Prev = list.tail
	list.tail = node
	return list
}

func (list *LinkedList[T]) RemoveLast() (*LinkedList[T], error ){
	if list.head == nil || list.tail == nil {
		return nil, errors.New("Can not remove from am empty list")
	}
	
	list.length -= 1
	if list.length == 0 {
		list.head = nil
		list.tail = nil
		return list, nil
	}
	beforeTail := list.tail.Prev
	beforeTail.Next = nil
	list.tail = beforeTail;

	return list, nil
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
