package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Let's learn web requests in golang.")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response:", res)
	fmt.Printf("Type of response:%T\n", res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data:", data)
}
