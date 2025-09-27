package main

import "fmt"

func main() {
	//声明map,key是string，value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1是一个空map")
	}

	myMap1 = make(map[string]string, 10)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap1)
}
