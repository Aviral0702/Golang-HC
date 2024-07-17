package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://api.github.com/users/Aviral0702?q=hello&name=aviral"

func main() {
	fmt.Println("Welcome to handle urls with me")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println("Result Scheme is: ", result.Scheme)
	fmt.Println("Result Host is: ", result.Host)
	fmt.Println("Result Path is: ", result.Path)

	qparams := result.Query()
	fmt.Printf("The type of query params is: %T\n", qparams)

	for _, values := range qparams {
		fmt.Println("Param is: ", values)
	}
}