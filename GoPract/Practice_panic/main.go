package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
func slicePanic() {
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}

func main() {
	slicePanic()
	fmt.Println("normally returned from main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
