package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num * 2
}

func main() {
	num := 2
	ptr := &num

	fmt.Println("Num has value:", num)
	fmt.Println("pointer contains:", ptr)
	fmt.Println("value of pointer:", *ptr)

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value contains:", value)
}
