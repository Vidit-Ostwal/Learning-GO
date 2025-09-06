package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	logincount := 23
	var result string

	if logincount < 10 {
		result = "Regular User"
	} else if logincount > 10 {
		result = "Power User"
	}

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

	fmt.Println(result)
}
