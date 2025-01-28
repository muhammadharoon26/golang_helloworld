package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
}

func sayHi() {
	fmt.Println("Hi, World!")
}

func main() {
	fmt.Println("Let's learn goroutine in golang.")

	go sayHello()
	sayHi()

	time.Sleep(1000 * time.Millisecond)
}
