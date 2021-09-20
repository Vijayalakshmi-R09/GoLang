package main

import (
	"fmt"
)

func main() {
	var letter string 
	fmt.Println("Enter a letter: ")
	fmt.Scanln(&letter)
	fmt.Printf("Letter %s is a ", letter)
	switch  {
	case letter=="a" || letter== "e" || letter== "i" || letter== "o" || letter== "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}
