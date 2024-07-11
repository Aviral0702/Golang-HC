package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please provide rating for pizza from 1 to 5")
	rating, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Give me rating +1 : ", numRating+1)
	}

}
