package main

import (
	"fmt"
	"testing"
)

// 结构体范性
type Node[T any] struct {
	Next  *Node[T]
	Value T
}

type Link[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *Link[T]) Append(t T) {
	node := &Node[T]{
		Next:  nil,
		Value: t,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *Link[T]) Print() {
	node := l.Head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

func TestGeneric(T *testing.T) {
	l := Link[int]{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Print()
}
