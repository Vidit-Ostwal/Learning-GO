package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	hitesh := User{"Hitest", "fsf", true, 1}
	fmt.Println(hitesh)
	fmt.Printf("hitest details are: %+v\n", hitesh)
	hitesh.getstatus()
	hitesh.NewMail()
	fmt.Printf("hitest details are: %+v\n", hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getstatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "NewEmail@gmail.com"
	fmt.Print("Your New EMail: ", u.Email)
}
