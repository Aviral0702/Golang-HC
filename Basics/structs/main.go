package main

import "fmt"

func main() {

	//There is no concept of inheritance and parent class in golang

	fmt.Println("Welcome to structs practice session")

	aviral := User{"Aviral", "aviral@gmail.com", true, 20}
	fmt.Println(aviral)
	fmt.Printf("Name: %+v\n", aviral)
}

type User struct {
	name   string
	email  string
	status bool
	age    int
}
