package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	var where = func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	fibParam := fib()
	where()
	for i := 1; i <= 10; i++ {
		fmt.Printf("fib at %d is %d\n", i, fibParam())
	}
	where()
}

func fib() func() int {
	current, next := 0, 1
	return func() int {
		current, next = next, current+next
		return current
	}
}
