package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Number is:", i)
	}

	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			break
		}
	}

	// numbers := []int{1, 2, 3, 4, 5}
	data := "Hello, world!"
	for index, char := range data {
		fmt.Printf("Index: %d, Value: %c\n", index, char)
	}
}
