package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	funcDouble := func(n int) int {
		return 2 * n
	}
	listDouble := mymap(funcDouble, list)
	fmt.Println(listDouble)
}

func mymap(f func(int) int, list []int) []int {
	res := make([]int, len(list))
	for i, num := range list {
		res[i] = f(num)
	}
	return res
}
