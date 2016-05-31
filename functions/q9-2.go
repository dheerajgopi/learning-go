package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	length int
	data   [2]int
}

func main() {
	var stck Stack
	sp := &stck
	sp.push(2)
	sp.push(3)
	fmt.Printf("%v", sp)
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

func (s *Stack) String() string {
	var str string
	for i := 0; i < s.length; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" +
			strconv.Itoa(s.data[i]) +
			"] "
	}
	return str + "\n"
}
