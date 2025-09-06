package main

import "fmt"

func main() {
	fmt.Println("Welcome to the maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("List of all languages: ", languages["JS"])

	// delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// Loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
