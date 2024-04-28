package main

import (
	"container/list"
	"fmt"
)

func main() {
	testList := list.New()
	testList.PushBack("1")
	testList.PushBack("2")
	testList.PushBack("3")

	for e := testList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
