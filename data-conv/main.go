package main

import "fmt"

func main() {
	fmt.Println("Let's learn data conversion in golang")

	num := 42
	fmt.Println("The value of num is:", num)
	fmt.Printf("Type of variable num is %T\n", num)

	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("The value of data is:", data)
	fmt.Printf("Type of variable num is %T\n", data)
}
