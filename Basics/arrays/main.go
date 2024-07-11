package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"

	fmt.Println("fruit list is : ", fruits)
	fmt.Println("Length of fruit list is : ", len(fruits))
}
