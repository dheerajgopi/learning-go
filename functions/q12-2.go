package main

import (
	"fmt"
	"strings"
)

func main() {
	list := []string{"a", "b", "c", "d", "e"}
	listUpperCase := mymap(strings.ToUpper, list)
	fmt.Println(listUpperCase)
}

func mymap(f func(string) string, list []string) []string {
	res := make([]string, len(list))
	for i, str := range list {
		res[i] = f(str)
	}
	return res
}
