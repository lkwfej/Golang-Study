package main

import "fmt"

func fool1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---foo3----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	r1 = 1000
	r2 = 2000
	return
}

func main() {
	c := fool1("abc", 999)

	fmt.Println("c = ", c)

	var ret1, ret2 int = fool3("foo", 333)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

}
