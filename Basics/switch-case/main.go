package main

import (
	"math/rand"
	"time"
)

func main() {
	// dice game
	// 1 to 6
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		println("You got one")
	case 2:
		println("You got two")
	case 3:
		println("You got three")
		fallthrough
	case 4:
		println("You got four")
	case 5:
		println("You got five")
	case 6:
		println("You got six")
	}

}
