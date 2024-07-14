package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["RUBY"] = "Ruby"
	fmt.Println(languages)

	//deletion in maps
	delete(languages, "RUBY")
	fmt.Println(languages)

	//Iterating over maps\
	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf(" Value: %v\n", value)
	}
}
