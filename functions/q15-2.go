package main

import "fmt"

func main() {
	addOne := plusN(1)
	fmt.Println(addOne(2))
}

func plusN(n int) func(int) int {
	return func(i int) int { return i + n }
}
