package main

import (
	"fmt"
)

func main() {
	arr := [50]int{}
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < 50; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	for i := 0; i < 50; i++ {
		fmt.Println("fib at ", i, "is ", arr[i])
	}

	arr2 := [10][10]int{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			arr2[i][j] = i*10 + j
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Println(arr2[i])
	}

	s := make([]byte, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = s[2:4]
	fmt.Println(len(s))
	fmt.Println(cap(s))

	items := [...]int{10, 20, 30, 40, 50}

	for index, _ := range items {
		items[index] *= 2
	}
	fmt.Println(items)

}
