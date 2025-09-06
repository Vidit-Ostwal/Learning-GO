package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	hitesh := User{"Hitest", "fsf", true, 1}
	fmt.Println(hitesh)
	fmt.Printf("hitest details are: %+v\n", hitesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
