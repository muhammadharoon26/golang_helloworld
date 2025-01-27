package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var haroon Person
	fmt.Println("Person:", haroon)
	haroon.FirstName = "m"
	haroon.LastName = "haroon"
	haroon.Age = 20
	fmt.Println("Person:", haroon)
}
