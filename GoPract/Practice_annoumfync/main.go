package main

import "fmt"

type add func(a int, b int) int

func main() {
	a := func() {
		fmt.Print("hello 0")
	}
	a()
	fmt.Printf("%T", a)
	func() {
		fmt.Println("\nhello world first class function")
	}()
	func(n string) {
		fmt.Println("Welcome", n)
	}("Back Squids")

	var a1 add = func(a int, b int) int {
		return a + b
	}
	s := a1(5, 6)
	fmt.Println("Sum", s)
}
