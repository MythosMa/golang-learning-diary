package main

import (
	"fmt"
	"reflect"

	"mythosma.com/struct/pack"
)

type TestStruct struct {
	a int     "test a"
	b float64 "test b"
	c string  "test c"
}

func main() {
	st := new(TestStruct)
	st.a = 10
	st.b = 5.6
	st.c = "hello"
	fmt.Printf("st = %+v\n", st)

	var st2 TestStruct = TestStruct{100, 7.1, "world"}
	fmt.Printf("st2 = %+v\n", st2)
	for i := 0; i < 3; i++ {
		refTag(st2, i)
	}

	st3 := new(pack.MyStruct)
	st3.Name = "hello"
	st3.Age = 10
	fmt.Printf("st3 = %+v\n", st3)
}

func refTag(tt TestStruct, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
