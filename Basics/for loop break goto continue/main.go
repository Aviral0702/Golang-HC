package main

import "fmt"

func main() {
	var days = []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for key, value := range days {
		if key == 3 {
			goto gonto
		}
		fmt.Printf("Key: %v, Value: %v\n", key+1, value)
	}

	//

gonto:
	fmt.Println("I am in goto statement")

}
