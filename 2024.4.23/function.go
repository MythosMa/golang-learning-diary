package main

import (
	"fmt"
)

type CompanyOptions struct {
	name    string
	address string
}

type SchoolOptions struct {
	name   string
	object string
}

type HomeOptions struct {
	city    string
	address string
}

type UserInfo struct {
	Company *CompanyOptions
	School  *SchoolOptions
	Home    *HomeOptions
}

func main() {
	value1, value2, value3 := calcFun(5, 5)
	fmt.Printf("和: %d, 乘积: %d, 差: %d\n", value1, value2, value3)
	value4, value5, value6 := calcFun2(5, 5)
	fmt.Printf("和: %d, 乘积: %d, 差: %d\n", value4, value5, value6)

	user := UserInfo{
		Company: &CompanyOptions{
			name:    "ABC",
			address: "Beijing",
		},
		School: &SchoolOptions{
			name:   "DDD",
			object: "English",
		},
		Home: &HomeOptions{
			city:    "Shanghai",
			address: "Changan",
		},
	}
	printUserInfo(user)

	unknowFunc(1, 100, "abccc", 41.1457, "ere", false)

	deferTest()

	fmt.Println(f())
}

func f() (ret int) {
	defer func() {
		ret += 10
	}()
	return 1
}

func calcFun(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func calcFun2(a, b int) (sum, product, difference int) {
	sum = a + b
	product = a * b
	difference = a - b
	return
}

func printUserInfo(user UserInfo) {
	if user.Company != nil {
		fmt.Printf("Company name: %s, address: %s\n", user.Company.name, user.Company.address)
	}
	if user.School != nil {
		fmt.Printf("School name: %s, object: %s\n", user.School.name, user.School.object)
	}
	if user.Home != nil {
		fmt.Printf("Home city: %s, address: %s\n", user.Home.city, user.Home.address)
	}
}

func unknowFunc(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Printf("int: %d\n", v)
		case string:
			fmt.Printf("string: %s\n", v)
		case float64:
			fmt.Printf("float64: %f\n", v)
		default:
			fmt.Println("unknown type: ", v)
		}
	}
}

func deferTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
