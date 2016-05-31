package main

import "fmt"

func main() {
	printVarArgs(1, 2, 3, 4)
}

func printVarArgs(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}
