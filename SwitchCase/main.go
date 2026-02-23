package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case in go lang")

	// No explicit seeding needed in Go 1.20+
	diceNumber := rand.Intn(6) + 1 // Generates a random number between 0 and 5, then adds 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open")
    case 2:
		fmt.Println("You can move 2 spot")
		// fallthrought moves to next place also cases
		// It moves to next and next moves
    case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
    case 4:
		fmt.Println("You can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
    case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("You can move to 5 spot")		 							
	}
}
