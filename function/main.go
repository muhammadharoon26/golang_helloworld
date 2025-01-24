package main

import "fmt"

func simpleFnction() {
	fmt.Println("simple function")
}

func add(a, b int) int {
	fmt.Print(a, " + ", b, " = ")
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b

	return
}

func main() {
	fmt.Println("we are learning function in Golang")
	simpleFnction()
	fmt.Println("Let's do some addition!,")
	fmt.Println(add(2, 3))
	fmt.Println("Let's do some multiplication!,")
	fmt.Println(multiply(2, 3))
}
