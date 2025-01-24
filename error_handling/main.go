package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("started error handling in the functions.")
	ans, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error Handling")
	}
	fmt.Println("The answer of division is:", ans)
}
