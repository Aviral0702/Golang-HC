package main

import "fmt"

func main() {
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proAd, proMsg := proAdder(2, 3, 5, 4, 6)
	fmt.Println("ProAdder result is: ", proAd)
	fmt.Println("ProAdder message is: ", proMsg)
}

func greeter() {
	fmt.Println("Hello welcome to functions class")
}

func adder(valA int, valB int) int {
	return valA + valB
}

func proAdder(values ...int) (int, string) {
	result := 0
	for _, value := range values {
		result += value
	}
	return result, "All values are added successfully"
}
