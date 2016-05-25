package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "asSASA ddd dsjkdsjs dk"
	s_bytes := []byte(s)
	fmt.Printf("Number of characters: %d\nNumber of bytes: %d\n", len(s_bytes), utf8.RuneCount(s_bytes))
}
