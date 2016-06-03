package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Calculate: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	var token string
	stck := new(stack.Stack)
	sp := &stck
	for _, char := range input {
		switch {
		case char == '+':
			sp.Push(sp.Pop() + sp.Pop())

		case char == '*':
			sp.Push(sp.Pop() * sp.Pop())

		case char == '-':
			sp.Push(sp.Pop() - sp.Pop())

		case char == '/':
			sp.Push(sp.Pop() / sp.Pop())

		case char == ' ':
			if token != "" {
				num, err := strconv.Atoi(token)
				if err != nil {
					fmt.Println("Invalid input")
					return
				}
				sp.Push(num)
				token = ""
			}

		default:
			token = token + string(char)
		}
	}
	if stck.Height == 1 {
		fmt.Printf("The result is : %d\n", stck.Pop())
		return
	}
	fmt.Println("Invalid expression!")
}
