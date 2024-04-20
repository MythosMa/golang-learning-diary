package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("hello world!")

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	var d, e, f string = "hello", " ", "world!"
	fmt.Println(d + e + f)

	g := 1
	// h := "+"
	i := 1
	fmt.Println(g + i)
	// fmt.Println(g + h + i)

	j := 1.0
	k := &j

	fmt.Println(j, k, *k)

	var aa, bb int
	var cc string
	aa, bb, cc = 1, 2, "hello"
	fmt.Println(aa, bb, cc)

	const (
		aaa = iota
		bbb
		ccc
		ddd
		eee = 1
		fff
		ggg
		hhh = "a"
		iii
		jjj
		kkk = "b"
		lll
		mmm
		nnn
		ooo = iota
	)

	fmt.Println(aaa, bbb, ccc, ddd, eee, fff, ggg, hhh, iii, jjj, kkk, lll, mmm, nnn, ooo)

	var str string = "hello, will"
	fmt.Println(strings.Replace(str, "ll", "ff", 3))

	fmt.Println(strings.Repeat("abc", 12))
	fmt.Println(strings.Split("abc,def,ghi", ","))
	fmt.Println(strings.Fields("abc def    ghi"))

}