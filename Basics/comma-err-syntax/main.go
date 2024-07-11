package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user inpur"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
}
