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
}
