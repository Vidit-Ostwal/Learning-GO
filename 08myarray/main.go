package main

import "fmt"

func main() {
	fmt.Println("Welcome to array of GoLang")

	var fruitList [4]string

	fruitList[0] = "vidit"
	fruitList[1] = "is"
	fruitList[3] = "best"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

}
