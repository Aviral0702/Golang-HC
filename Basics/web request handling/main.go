package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const url = "https://api.github.com/users/Aviral0702"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred %T\n", err)
	}
	defer resp.Body.Close()

	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
