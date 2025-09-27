package main

import "fmt"

func m_ain() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine正在运行")

		c <- 666
	}()

	num := <-c

	fmt.Println("num = ", num)
	fmt.Println("main goroutine结束...")

}
