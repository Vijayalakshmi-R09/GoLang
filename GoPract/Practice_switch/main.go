package main

import (
	"fmt"
)

func main() {
	var finger int
	var colour int
	fmt.Println("Enter the finger number: ")
	fmt.Scanln(&finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Invalid entry")
	}
	colour = finger
	fmt.Printf("Matching Colors to finger %d: ",finger)
	switch colour {
	case 1:
		fmt.Println(" Green")
	case 2:
		fmt.Println(" Red")
	case 3:
		fmt.Println(" Grey")
	case 4:
		fmt.Println(" Gold")
	case 5:
		fmt.Println(" Pink")
	default:
		fmt.Println("No Color")
	}
}
