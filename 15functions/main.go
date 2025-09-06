package main

import "fmt"

func main() {
	greater()
	fmt.Println("Welcome to functions in GoLang")
	greater()
	greater2()

	// Function inside a function is not allowed
	// result := adder(3, 5)
	result := proAdder(4, 4, 5, 667, 7, 5, 3, 3, 4566543, 1)
	fmt.Println("Result is: ", result)

}

func greater2() {
	fmt.Println("Another function")
}

func adder(var1 int, var2 int) int {
	return var1 + var2
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greater() {
	fmt.Println("Namasthe from GoLang")
}
