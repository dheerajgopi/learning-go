package main

import (
	"fmt"
)

func main() {
	s := "asSASA ddd dsjkdsjs dk"
	s = s[:4] + "abc" + s[7:]
	fmt.Println(s)
}
