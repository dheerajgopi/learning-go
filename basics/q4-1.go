package main

import "fmt"

func main() {
	a := "A"
	for i := 0; i < 100; i++ {
		fmt.Println(a)
		a = a + "A"
	}
}
