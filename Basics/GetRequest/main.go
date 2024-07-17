package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request program")
	// performGetReq()
	// performPostReq()
	performPostFormReq()
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

func performPostFormReq() {
	const myurl = "http://localhost:8000/postform"

	//fake form payload
	payload := url.Values{}

	payload.Add("firstname", "Aviral")
	payload.Add("lastname", "Asthana")
	payload.Add("email", "aviral@go.dev")

	response, err := http.PostForm(myurl, payload)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response: ", string(content))

}
