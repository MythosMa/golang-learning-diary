package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	originArr := randomArray()
	arr1 := originArr[:]
	arr2 := originArr[:]
	startTime := time.Now()
	bubbleSort(arr1)
	elapsedTime := time.Since(startTime)
	fmt.Println("Bubble Sort Time:", elapsedTime)
	fmt.Println("After Bubble Sort:", arr1)
	startTime = time.Now()
	quickSort(arr2)
	elapsedTime = time.Since(startTime)
	fmt.Println("Quick Sort Time:", elapsedTime)
	fmt.Println("After Quick Sort:", arr2)
}

func randomArray() []int {
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = i + 1
	}
	for i := len(arr) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]

	left := make([]int, 0)
	right := make([]int, 0)
	for i, v := range arr {
		if i == len(arr)-1 {
			break
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	quickSort(left)
	quickSort(right)

	result := append(left, pivot)
	result = append(result, right...)
	copy(arr, result)
}
