package main

import "fmt"

type UserInfo struct {
    name string
	age int
	sex string
}

func main() {
    user_1 := UserInfo{name: "张三", age: 18, sex: "男"}
	user_2 := UserInfo{name: "李四", age: 20, sex: "女"}
	fmt.Println(user_1, user_2)

	var user_3 UserInfo
	user_3.name = "王五"
	user_3.age = 22
	user_3.sex = "男"
	fmt.Printf("name: %s, age: %d, sex: %s\n", user_3.name, user_3.age, user_3.sex)

	var use_4 UserInfo
	use_4 = UserInfo{"赵六", 24, "男"}
	printUserInfo(use_4)
	fmt.Printf("in main => name: %s, age: %d, sex: %s\n", use_4.name, use_4.age, use_4.sex)
	printUserInfoByPoint(&use_4)
	fmt.Printf("in main after point => name: %s, age: %d, sex: %s\n", use_4.name, use_4.age, use_4.sex)
}

func printUserInfo(user UserInfo) {
	user.name = "小明"
	fmt.Printf("in func => name: %s, age: %d, sex: %s\n", user.name, user.age, user.sex)
}

func printUserInfoByPoint(user *UserInfo) {
	user.name = "小红"
	fmt.Printf("in func point => name: %s, age: %d, sex: %s\n", user.name, user.age, user.sex)
}