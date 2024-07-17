package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request program")
	// performGetReq()
	performPostReq()
}

func performGetReq() {
	const url = "http://localhost:8000/get"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Response content-length : ", resp.ContentLength)
	var responseString strings.Builder

	response, _ := io.ReadAll(resp.Body)
	// content := string(response)
	byteCount, _ := responseString.Write(response)
	fmt.Println("Byte count: ", byteCount)
	fmt.Println("Response is here: ", responseString.String())
	// fmt.Println("Response is here: ", content)
}

func performPostReq() {
	const url = "http://localhost:8000/post"

	//fake json payload
	payload := strings.NewReader(`
		{
			"coursename":"golang",
			"price":0,
			"platform":"udemy"
		}
	`)
	resp, err := http.Post(url, "application/json", payload)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)
	fmt.Println("Response is here: ", string(content))
}
