package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aviral0702/mongoApi/router"
)

func main() {
	fmt.Println("Server is running now...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))

}
