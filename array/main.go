package main

import "fmt"

func main() {
	fmt.Println("we are learning array in golang.")

	var name [5]string
	name[0] = "prince"
	name[2] = "haroon"
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Names of person is:", name)
	fmt.Println("Numbers are:", numbers)
	fmt.Println("Length of Numbers array is:", len(numbers))
	fmt.Println("Length of Name array is:", len(name))
	fmt.Println("value of name at 2 index is:", name[2])
	fmt.Println("----------------------------------------")
	fmt.Println("----------------------------------------")

	var price [5]int
	fmt.Println("price is:", price)
}
