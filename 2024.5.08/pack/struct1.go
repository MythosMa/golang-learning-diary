package pack

// 只有大写才能被其他文件使用，如果小写，则只能在本文件中使用，例如 address string，在main中就无法被调用到
type MyStruct struct {
	Name    string
	Age     int
	address string
}
