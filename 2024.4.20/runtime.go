package main


import (
	"fmt"
	"runtime"
	"os"
)

func main() {
    var goos = runtime.GOOS
	fmt.Println(goos)
	var path = os.Getenv("PATH")
	fmt.Println(path)
}