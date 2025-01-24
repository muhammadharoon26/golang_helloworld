package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array: ", numbers)
	fmt.Println("Length:", len(numbers))
	numbers = append(numbers, 6, 7, 8, 9, 10)
	fmt.Println("Numbers array: ", numbers)
	fmt.Println("Length:", len(numbers))

	name := []string{}
	fmt.Println("Name array: ", name)
}
