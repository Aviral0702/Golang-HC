package main

import "fmt"

func main() {
	fmt.Println("Welcome to ifelse")

	// if else
	age := 20
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	if num := 3; num < 4 {
		fmt.Println("Number is less than 4")
	} else {
		fmt.Println("Number is greater than 4")
	}
}
