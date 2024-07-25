package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //this is a steroid version of time.millisecond

func main() {
	// go greeter("hello")
	// greeter("world")
	websiteList := []string{
		"https://google.com",
		"https://fb.com",
		"https://amazon.com",
		"https://youtube.com",
	}

	for _, web := range websiteList {
		// time.Sleep(3 * time.Millisecond) //this doesn't work as expected I tried
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(msg string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(msg)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
