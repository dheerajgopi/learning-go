package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value      int
	Prev, next *Node
}

type List struct {
	head, tail *Node
}

func (l *List) Front() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) Push(v int) *List {
	n := &Node{Value: v}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.Prev = l.tail
	}
	l.tail = n

	return l
}

func (l *List) Pop() (v int, err error) {
	if l.tail == nil {
		err = errors.New("List is empty!")
	} else {
		v = l.tail.Value
		l.tail = l.tail.Prev

		if l.tail == nil {
			l.head = nil
		}
	}

	return v, err
}

func main() {
	l := new(List)
	l.Push(1)
	l.Push(2)
	l.Push(3)

	for val, err := l.Pop(); err == nil; val, err = l.Pop() {
		fmt.Println(val)
	}
}
