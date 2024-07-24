package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("hello")
	greeter("world")
}

func greeter(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(msg)
	}
}
