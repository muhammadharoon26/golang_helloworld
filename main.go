package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello World!")
	myutil.PrintMessage("Hello World")

	var name string = "Haroon"
	var version = "2.0"
	var money int = 200000
	var currency = 200000

	fmt.Println("My name is", name, "!")
	fmt.Println("Its", version)
	fmt.Println("money:", money)
	fmt.Println("currency:", currency)
}
