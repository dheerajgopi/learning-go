package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: cat [filename]")
		return
	}

	filename := args[0]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File is corrupted or does not exist")
	}

	fmt.Printf("%s", string(content))
	return
}
