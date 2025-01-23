package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello World!")
	myutil.PrintMessage("Hello World")

	var name string = "My name is Haroon!"
	var version = "2.0"
	var money int = 200000
	var currency = 200000

	fmt.Println(name)
	fmt.Println(version)
	fmt.Println(money)
	fmt.Println("currency: ", currency)
}
