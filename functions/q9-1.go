package main

import "fmt"

type Stack struct {
	length int
	data   [2]int
}

func main() {
	var stck Stack
	sp := &stck
	sp.push(2)
	fmt.Println(sp.pop())
	sp.pop()
	fmt.Println(sp.length)
	fmt.Println(sp.data)
}

func (s *Stack) push(n int) {
	if s.length >= len(s.data) {
		return
	} else {
		s.data[s.length] = n
		s.length++
	}
}

func (s *Stack) pop() int {
	if s.length <= 0 {
		fmt.Println("Stack is empty")
		return 0
	} else {
		s.length--
		return s.data[s.length]
	}
}
