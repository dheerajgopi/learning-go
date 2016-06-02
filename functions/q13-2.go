package main

import "fmt"

func main() {
	lst := []int{11, 2, 5, 6}
	fmt.Println(min(lst))
}

func min(list []int) (minNum int) {
	minNum = list[0]
	for _, num := range list {
		if num < minNum {
			minNum = num
		}
	}
	return
}
