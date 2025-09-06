package main

import "fmt"

func main() {
	fmt.Println("Hello Vidit")
	var username string = "Vidit"
	fmt.Printf("Variable is of type: %T \n", username)

	fmt.Println("Hello Vidit")
	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	fmt.Println("Hello Vidit")
	var smallval int = 256
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallfloat float32 = 253.35645766453
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

}
