package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Lets's learn strings in golang")

	data := "apple, orange, banana"
	fmt.Println(data)
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four five two"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str = "            Hello, World!             "
	fmt.Println(str)
	fmt.Println(strings.TrimSpace(str))
}
