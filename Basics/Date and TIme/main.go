package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study with golang")
	presentTime := time.Now()
	fmt.Println("Current date is: ", presentTime.Format("02-01-2006"))
	fmt.Println("Current time is: ", presentTime.Format("15:04:05"))

	createdDate := time.Date(2024, time.February, 07, 04, 55, 0, 0, time.UTC)
	fmt.Println("Created date is: ", createdDate.Format("02-01-2006 15:04:05"))

}
