package main

import "fmt"

type Stack struct {
	i    int
	data [10]int
}

func main() {
	var stck Stack
	sp := &stck
	sp.push(2)
	fmt.Println(sp.pop())
	sp.pop()
	fmt.Println(sp.i)
	fmt.Println(sp.data)
}

func (s *Stack) push(n int) {
	if s.i >= len(s.data) {
		return
	} else {
		s.data[s.i] = n
		s.i++
	}
}

func (s *Stack) pop() int {
	if s.i <= 0 {
		fmt.Println("Stack is empty")
		return 0
	} else {
		s.i--
		return s.data[s.i]
	}
}
