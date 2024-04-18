package main

import (
	"fmt"
)

func print(str string) {
	for i := 0 ; i < 10 ; i++ {
		fmt.Printf("%s=>%d\n", str, i + 1)
	}
}

func calc(a int, b int, c chan int) {
	c <- (a + b)
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
	close(c)
}

func main() {
    // go print("a")
	// go print("b")
	// print("c")


	// c := make(chan int)
	// go calc(1, 2, c)
	// go calc(3, 4, c)
	// r1, r2 := <-c, <-c
	// fmt.Println(r1, r2)


	// d := make(chan int, 1)
	// fmt.Println("11111111111111111")
	// d <- 1
	// fmt.Println("22222222222222222")
	// d <- 2
	// fmt.Println("33333333333333333")
	// fmt.Println(<-d)
	// fmt.Println("444444444444444444")
	// fmt.Println(<-d)

	a := make(chan int, 10)
	go fibonacci(cap(a), a)
	for i := range a {
	    fmt.Println(i)
	}
}


