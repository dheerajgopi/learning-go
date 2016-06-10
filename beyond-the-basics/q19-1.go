package main

import (
	"fmt"
)

type e interface{}

func double(n e) e {
	switch n.(type) {
	case int:
		return 2 * n.(int)

	case string:
		return n.(string) + n.(string)
	}
	return n
}

func mymap(f func(e) e, list []e) []e {
	res := make([]e, len(list))
	for i, elem := range list {
		res[i] = f(elem)
	}
	return res
}

func main() {
	listInt := []e{1, 2, 3, 4, 5}
	listStr := []e{"a", "b", "c"}
	listDoubleInt := mymap(double, listInt)
	listDoubleStr := mymap(double, listStr)
	fmt.Println(listDoubleInt)
	fmt.Println(listDoubleStr)
}
