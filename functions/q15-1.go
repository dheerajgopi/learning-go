package main

import "fmt"

func main() {
	addTwo := plusTwo()
	fmt.Println(addTwo(2))
}

func plusTwo() func(int) int {
	return func(n int) int { return n + 2 }
}
