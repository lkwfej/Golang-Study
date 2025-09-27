package main

import (
	"fmt"
)

func main() {
	var a int //默认值是0
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Println("bb = %s,type of bb = %T\n", bb, bb)

	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd:"
	fmt.Println("cc = %s,type of cc = %T\n", cc, cc)

	e := 100
	fmt.Println("e = ", e)

	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ",yy = ", yy)

}
