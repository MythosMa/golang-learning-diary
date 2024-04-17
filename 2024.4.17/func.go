package main

import "fmt"

func main() {
	a, b := 1, 2
	swapNum(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swapNumByReturn(a, b))

	c, d := "hello", "world"
	swapStr(&c, &d)
	fmt.Println(c, d)
}

func swapNum(x *int, y *int) {
    *x, *y = *y, *x
}

func swapNumByReturn(x, y int)(int, int){
	return y, x
}

func swapStr(x *string, y *string) {
	*x, *y = *y, *x
}