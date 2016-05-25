package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var flag bool

		if i%3 == 0 {
			fmt.Printf("Fizz")
			flag = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			flag = true
		}
		if !flag {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
