package main

import "fmt"

const (
	BERJING = iota
	SHANGHAI
	SHENZHEN
)

func main() {
	//常量
	const length int = 10

	fmt.Println("length = ", length)

	//length = 100 //常量是不允许修改的
}
