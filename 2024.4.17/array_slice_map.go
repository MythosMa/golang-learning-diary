package main
import "fmt"

func main() {
    // arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// fmt.Println("Original Array:", arr)
	// fmt.Println("Original Array len:", len(arr))
	// fmt.Println("Original Array cap:", cap(arr))
	// fmt.Println("Original Slice:", slice)
	// fmt.Println("Original Slice len:", len(slice))
	// fmt.Println("Original Slice cap:", cap(slice))

	var nums []int
	nums = append(nums, 0)
	customPrint(nums)
	nums = append(nums, 1)
	customPrint(nums)
	nums = append(nums, 2, 3, 4, 5, 6)
	customPrint(nums)

	var slice = make([]int, 10, 20)
	customPrint(slice)

	var slice_2 []int
	fmt.Println(slice_2)
	fmt.Println(slice_2 == nil)

	for index, value := range nums {
		fmt.Printf("index=%d, value=%d\n", index, value)
	}

	for index, value := range "hello world" {
		fmt.Printf("index=%d, value=%c\n", index, value)
	}

	var assetsCountInfo = make(map[string]int, 0)
	assetsCountInfo["picture"] = 10
	assetsCountInfo["video"] = 20
	assetsCountInfo["audio"] = 30
	fmt.Println(assetsCountInfo)
	fmt.Println(assetsCountInfo["picture"])
	value, isOk := assetsCountInfo["picture1"]
	fmt.Printf("value: %d, isOk: %t\n", value, isOk)
	for key, value := range assetsCountInfo {
	    fmt.Printf("key=%s, value=%d\n", key, value)
	}
}

func customPrint(x []int){
	fmt.Printf("len=%d, cap=%d, value=%v\n", len(x), cap(x), x)
}