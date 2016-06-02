package main

import "fmt"

func main() {
	lst := []int{11, 2, 5, 6}
	fmt.Println(max(lst))
}

func max(list []int) (maxNum int) {
	maxNum = list[0]
	for _, num := range list {
		if num > maxNum {
			maxNum = num
		}
	}
	return
}
