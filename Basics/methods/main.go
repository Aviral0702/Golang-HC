package main

import "fmt"

func main() {

	//There is no concept of inheritance and parent class in golang

	fmt.Println("Welcome to methods practice session")

	aviral := User{"Aviral", "aviral@gmail.com", true, 20}
	fmt.Println(aviral)
	fmt.Printf("Name: %+v\n", aviral)
	aviral.GetStatus()
	aviral.NewMail()
	fmt.Println(aviral)
	aviral.NewMailWithPointer()
	fmt.Println(aviral)

}

type User struct {
	name   string
	email  string
	active bool
	age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.active)
}

func (u User) NewMail() {
	u.email = "aviral@go.dev"
	fmt.Println("New email is: ", u.email)
}

func (u *User) NewMailWithPointer() {
	u.email = "aviral@go.dev"
	fmt.Println("New email is: ", u.email)
}
