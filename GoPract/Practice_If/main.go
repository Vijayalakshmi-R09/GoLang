package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter Number: ")
	var num int
	fmt.Scanln(&num)

	if n := num; n%2 == 0 {
		fmt.Println("The number", n, "is even")
	} else {
		fmt.Println("The number ", n, " is odd")
	}

	//fmt.Println(n) n is accesible only within if and else not outside

}
