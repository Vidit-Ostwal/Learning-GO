package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://go.dev/"

func main() {
	fmt.Println("Welcome to the Handling URL in Golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
}
