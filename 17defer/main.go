package main

import "fmt"

func main() {
	// fmt.Println("Welcome to defer statement in GoLang 1")

	defer fmt.Println("Welcome to defer statement in GoLang 3")
	defer fmt.Println("Welcome to defer statement in Golang 2")

	fmt.Println("Welcome to defer statement in GoLang 1")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
