package main

import (
	"container/list"
	"fmt"
)

func main() {
	linkList := list.New()
	linkList.PushBack(1)
	linkList.PushBack(2)
	linkList.PushBack(4)

	for elem := linkList.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value)
	}
}
