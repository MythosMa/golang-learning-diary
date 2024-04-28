package main

import "fmt"

func main() {
	fmt.Println("7-14 answer:", exercise7_14("Google"))
	exercise7_15()

	x10Func := func(data int) int {
		return data * 10
	}
	fmt.Println("7-16 answer:", exercise7_16(x10Func, []int{1, 2, 3, 4, 5}))
}

func exercise7_14(s string) string {
	bytes := []byte(s)
	leftIndex, rightIndex := 0, len(bytes)-1
	for leftIndex < rightIndex {
		bytes[leftIndex], bytes[rightIndex] = bytes[rightIndex], bytes[leftIndex]
		leftIndex++
		rightIndex--
	}
	return string(bytes)
}

func exercise7_15() {
	ori := []rune{'a', 'b', 'c', 'c', 'd', 'd', 'd', 'e', 'f', 'g'}
	result := []rune{}
	result = append(result, ori[0])
	for i := 1; i < len(ori); i++ {
		if ori[i] != ori[i-1] {
			result = append(result, ori[i])
		}
	}
	fmt.Println(string(result))
}

type X10Function func(int) int

func exercise7_16(f X10Function, datas []int) []int {
	result := make([]int, len((datas)))
	for index, data := range datas {
		result[index] = f(data)
	}
	return result
}
