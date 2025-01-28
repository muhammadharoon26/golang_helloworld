package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Let's learn url handling!")
	myurl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Println("URL:", myurl)
	fmt.Printf("Type of url:%T\n", myurl)
	parsedurl, err := url.Parse(myurl)
	// Check if there was an error parsing the URL
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	// Print the parsed URL components
	fmt.Printf("Type of parsed url:%T\n", parsedurl)
	fmt.Println("Parsed url:", parsedurl)
	fmt.Println("Scheme:", parsedurl.Scheme)
	fmt.Println("Opaque:", parsedurl.Opaque)
	fmt.Println("User:", parsedurl.User)
	fmt.Println("Host:", parsedurl.Host)
	fmt.Println("Path:", parsedurl.Path)
	fmt.Println("RawQuery:", parsedurl.RawQuery)
	fmt.Println("Fragment:", parsedurl.Fragment)
	// You can also use the following methods to get the query parameters
	// as a map
	queryparams := parsedurl.Query()
	fmt.Println("Query parameters:", queryparams)
}
