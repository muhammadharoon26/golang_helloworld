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
	var dimension float64 = 101.101
	var decided bool = false
	decided = true

	fmt.Println("My name is", name, "!")
	fmt.Println("Its", version)
	fmt.Println("money:", money)
	fmt.Println("currency:", currency)
	fmt.Println(dimension)
	fmt.Println(decided)

	const pi = 67.12
	fmt.Println("pi value:", pi)

	person := 123
	fmt.Println(person)

	var Public = "data is important"
	var private = "data is private"

	fmt.Println(Public)
	fmt.Println(private)
}
