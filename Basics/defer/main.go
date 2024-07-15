package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("1")
	defer fmt.Println("2")
	myDefer()
}

func myDefer() {
	var data = []int{1, 2, 3, 4, 5}
	for value := range data {
		defer fmt.Println(value)
	}
}
