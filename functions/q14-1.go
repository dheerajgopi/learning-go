package main

import "fmt"

func main() {
	lst := []int{3, 2, 1, 1, 0, -1, 5, 7, 9}
	bubbleSort(lst)
	fmt.Println(lst)
}

func bubbleSort(list []int) {
	flag := false
	for !flag {
		flag = true
		for i, num := range list {
			if i < len(list)-1 && num > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				flag = false
			}
		}
	}
}
